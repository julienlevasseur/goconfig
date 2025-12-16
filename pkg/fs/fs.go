package fs

import (
	"io/fs"
	"os"
)

func Chmod(file string, mode fs.FileMode) {
	os.Chmod(file, mode)
}

func Chown(file string, uid, gid int) {
	os.Chown(file, uid, gid)
}
