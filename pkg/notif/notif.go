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
	t := true

	path1Sum, err := checksum.MD5sum(path1)
	if err != nil {
		return &t
	}

	path2Sum, err := checksum.MD5sum(path2)
	if err != nil {
		return &t
	}

	if path1Sum == path2Sum {
		return &t
	}

	f := false
	return &f
}

func NotSameFile(path1, path2 string) *bool {
	f := false

	path1Sum, err := checksum.MD5sum(path1)
	if err != nil {
		return &f
	}

	path2Sum, err := checksum.MD5sum(path2)
	if err != nil {
		return &f
	}

	if path1Sum == path2Sum {
		return &f
	}

	t := true
	return &t
}
