package notif

import (
	"fmt"
	"os"

	"github.com/codingsince1985/checksum"
)

func IgnoreDueToNotIf(name, step string) {
	fmt.Printf("[%v][%v] Ignore due to NotIf\n", name, step)
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

func SameFile(path1, path2 string) *bool {
	if checksum.MD5Sum(path1) == checksum.MD5Sum(path2) {
		t := true
		return &t
	}

	f := false
	return &f
}
