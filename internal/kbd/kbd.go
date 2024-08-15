package kbd

import (
	"os"
	"strings"
)

func GetKeyboards() ([]string, error) {
	keyboards := []string{}
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
		deviceName = deviceName[1 : len(deviceName)-1]
		if strings.Contains(strings.ToLower(deviceName), "keyboard") {
			keyboards = append(keyboards, device)
		}
	}

	return keyboards, nil
}
