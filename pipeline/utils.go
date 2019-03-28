package pipeline

import (
	"os"
)


func IsInitialized(filepath string) (bool, error) {
	info, err := os.Stat(filepath)
	if err == nil && info.Mode().IsRegular() {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

