package notif

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/file"
)

func IgnoreDueToNotIf(name string) {
	fmt.Printf("[%v][Install] Ignore Installation due to NotIf\n", name)
}

func FileExists(path string) *bool {
	return file.Exists(path)
}
