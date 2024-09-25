package kbd

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"syscall"

	"github.com/halfloafhq/keymulate/internal/audio"
)

type inputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

func GetKeyboards() (map[string]string, error) {
	keyboards := make(map[string]string)
	inputFile, err := os.ReadFile("/proc/bus/input/devices")
	if err != nil {
		return nil, err
	}

	devices := strings.Split(string(inputFile), "\n\n")
	for _, device := range devices {
		deviceInfo := strings.Split(device, "\n")

		if len(deviceInfo) < 2 {
			continue
		}

		deviceN := deviceInfo[1]
		deviceName := strings.Split(deviceN, "N: Name=")[1]
		deviceName = deviceName[1:]

		if strings.Contains(strings.ToLower(deviceName), fmt.Sprintf("keyboard\"")) {

			deviceName = deviceName[:len(deviceName)-1]

			if _, exists := keyboards[deviceName]; !exists {
				keyboards[deviceName] = device
			}

		}

	}

	return keyboards, nil
}

func GetEvents(keyboards map[string]string) []string {
	events := []string{}

	for _, keyboard := range keyboards {
		deviceH := strings.Split(keyboard, "\n")[5]
		handlers := strings.Split(strings.Split(deviceH, "H: Handlers=")[1], " ")
		for _, handler := range handlers {
			if strings.Contains(handler, "event") {
				events = append(events, handler)
			}
		}
	}

	return events
}

func Listen(switchOpt string, events []string) error {

	if !isValidSwitch(switchOpt) {
		fmt.Printf("Invalid switch name entered. Please insert a valid switch name.\n")
		return errors.New("Invalid switch name")
	}

	var wg sync.WaitGroup

	otoCtx := audio.LoadAudioCtx()

	press, release := audio.LoadSoundsForKeyboard(switchOpt)

	for _, event := range events {
		wg.Add(1)
		go func(eventPath string) {
			defer wg.Done()

			file, err := os.Open(fmt.Sprintf("/dev/input/%s", eventPath))
			if err != nil {
				fmt.Printf("Error opening %s: %v\n", eventPath, err)
				return
			}
			defer file.Close()

			for {
				event := inputEvent{}
				err := binary.Read(file, binary.LittleEndian, &event)
				if err != nil {
					if err == syscall.EINVAL {
						fmt.Printf("Device %s disconnected or no longer available\n", eventPath)
						return
					}
					fmt.Printf("Error reading from %s: %v\n", eventPath, err)
					continue
				}

				// Play audio based on event type and code
				if event.Type == 1 { // EV_KEY events
					var sound []byte
					var soundKey string

					if event.Value == 1 { // Key press
						soundKey = audio.GetSoundKey(event.Code, true)
						sound = press[soundKey]
					} else if event.Value == 0 { // Key release
						soundKey = audio.GetSoundKey(event.Code, false)
						sound = release[soundKey]
					}

					if sound != nil {
						go audio.PlaySound(otoCtx, sound)
					}
				}

				//				fmt.Printf("Event from %s: Type: %d, Code: %d, Value: %d\n",
				//					eventPath, event.Type, event.Code, event.Value)
			}
		}(event)
	}

	wg.Wait()
	return nil
}

func isValidSwitch(switchName string) bool {
	switch switchName {
	case "alpaca":
		return true
	case "blackink":
		return true
	case "bluealps":
		return true
	case "boxnavy":
		return true
	case "buckling":
		return true
	case "cream":
		return true
	case "holypanda":
		return true
	case "mxblack":
		return true
	case "mxblue":
		return true
	case "mxbrown":
		return true
	case "redink":
		return true
	case "topre":
		return true
	case "turquoise":
		return true
	default:
		return false
	}
}
