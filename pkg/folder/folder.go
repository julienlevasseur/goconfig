package folder

import (
	"errors"
	"fmt"
	"os"
)

func Create(path string, notIf bool) error {
	if !notIf {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModeDir)
			if err != nil {
				return err
			}
		}
	} else {
		fmt.Printf("[%v][%v] Ignore due to NotIf\n", "Folder", "Create")
	}

	return nil
}

func CreateAll(path string, notIf bool) error {
	if !notIf {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(path, os.ModeDir)
			if err != nil {
				return err
			}
		}
	} else {
		fmt.Printf("[%v][%v] Ignore due to NotIf\n", "Folder", "Create")
	}

	return nil
}

func Exists(path string) *bool {
	t := true
	_, err := os.ReadDir(path)
	if os.IsNotExist(err) {
		f := false
		return &f
	} else if err != nil {
		fmt.Printf("[ERROR] [Folder][Create] %v", err)
		return &t
	}

	return &t
}
