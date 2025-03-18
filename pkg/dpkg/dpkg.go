package dpkg

import (
	"fmt"
	"strings"

	"github.com/julienlevasseur/goconfig/pkg/command"
	"github.com/julienlevasseur/goconfig/pkg/file"
	"github.com/julienlevasseur/goconfig/pkg/notif"
)

type Dpkg struct {
	Name       string
	ArchiveURL string
	NotIf      *bool
}

func (d Dpkg) Install() error {
	var f bool = false

	if d.NotIf == nil || d.NotIf == &f {
		fmt.Printf("\n[%v][Download archive]\n", d.Name)

		fileName := strings.Split(
			d.ArchiveURL,
			"/",
		)[len(strings.Split(d.ArchiveURL, "/"))-1]

		downloadedFilePath := "/tmp/" + fileName

		err := file.Download(
			d.ArchiveURL,
			downloadedFilePath,
		)
		if err != nil {
			return err
		}

		err = command.Exec(
			"dpkg",
			[]string{
				"-i",
				downloadedFilePath,
			},
		)
		if err != nil {
			return err
		}
	} else {
		notif.IgnoreDueToNotIf(d.Name)
	}

	return nil
}
