package notif

import (
	"fmt"
	"os"
)

func IgnoreDueToNotIf(name, step string) {
	fmt.Printf("[%v][%v] Ignore due to NotIf", name, step)
}

func FileExists(path string) *bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f := false
		return &f
	} else {
		t := true
		return &t
	}
}
