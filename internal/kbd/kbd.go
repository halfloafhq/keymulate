package kbd

import (
	"fmt"
	"os"
	"strings"
)

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
