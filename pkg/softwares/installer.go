package softwares

import (
	"fmt"
	"os"

	"github.com/julienlevasseur/goconfig/pkg/archive"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

type IInstaller interface {
	// Install(name, version, arch, platform string, notIf ...bool)
	InstallFromArchiveURL(name, archiveURL, archiveFilename, destFolder string, notIf ...bool)
	InstallFromURLBinaryURL(BinaryURL, destFolder string, notIf ...bool)
}

type Installer struct{}

func (i *Installer) InstallFromArchiveURL(name, version, archiveURL, archiveFilename, destFolder string, notIf ...bool) {
	if notIf[0] {
		fmt.Printf("[%v][Install] Ignore Install due to NotIf\n", name)
	} else {
		fmt.Printf("[%v][Install] Download archive\n", name)

		err := file.Download(
			archiveURL,
			archiveFilename,
		)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Decompress archive\n", name)
		err = archive.Unzip(archiveFilename, destFolder)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Delete archive\n", name)
		err = file.Delete(archiveFilename)
		if err != nil {
			panic(err)
		}

		err = os.Rename(
			fmt.Sprintf(
				"%v/%v",
				destFolder,
				name,
			),
			fmt.Sprintf(
				"%v/%v_%v",
				destFolder,
				name,
				version,
			),
		)
		if err != nil {
			panic(err)
		}

		err = os.Symlink(
			fmt.Sprintf(
				"%v/%v_%v",
				destFolder,
				name,
				version,
			),
			fmt.Sprintf(
				"%v/%v",
				destFolder,
				name,
			),
		)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Complete\n", name)
	}
}

func (i Installer) InstallFromURLBinaryURL(BinaryURL, destFolder string, notIf ...bool) {
	// TODO
}
