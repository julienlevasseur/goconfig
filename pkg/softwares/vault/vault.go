package vault

import (
	"fmt"
	"os"

	"github.com/julienlevasseur/goconfig/pkg/archive"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

const destFolder = "/usr/local/bin"

type VaultInstaller struct{}

func (vi *VaultInstaller) Install(name, version, arch, platform string, notIf ...bool) {
	if notIf[0] {
		fmt.Printf("[%v][Install] Ignore Install due to NotIf\n", name)
	} else {
		fmt.Printf("[%v][Install] Download archive\n", name)
		localFileName := fmt.Sprintf("vault_%v_%v_%v.zip", version, platform, arch)
		URL := fmt.Sprintf("https://releases.hashicorp.com/vault/%v/vault%v_%v_%v.zip", version, version, platform, arch)

		err := file.Download(
			URL,
			localFileName,
		)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Decompress archive\n", name)
		err = archive.Unzip(localFileName, destFolder)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Delete archive\n", name)
		err = file.Delete(localFileName)
		if err != nil {
			panic(err)
		}

		// Rename file to vault_${version}
		err = os.Rename(
			fmt.Sprintf(
				"%v/vault",
				destFolder,
			),
			fmt.Sprintf(
				"%v/vault_%v",
				destFolder,
				version,
			),
		)
		if err != nil {
			panic(err)
		}

		// Link vault_${version} to vault
		err = os.Symlink(
			fmt.Sprintf(
				"%v/vault_%v",
				destFolder,
				version,
			),
			fmt.Sprintf(
				"%v/vault",
				destFolder,
			),
		)
		if err != nil {
			panic(err)
		}

		// Delete LICENSE.txt file (from archive)
		err = file.Delete(fmt.Sprintf("%v/LICENSE.txt", destFolder))
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Complete\n", name)
	}
}
