package terraform

import (
	"fmt"
	"os"

	"github.com/julienlevasseur/goconfig/pkg/archive"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

const name = "Terraform"
const destFolder = "destFolder"

func Install(version, arch, platform string, notIf ...bool) {
	if notIf[0] {
		fmt.Printf("[%v][Install] Ignore Install due to NotIf\n", name)
	} else {
		fmt.Printf("[%v][Install] Download archive\n", name)
		localFileName := fmt.Sprintf("terraform_%v_%v_%v.zip", version, platform, arch)
		URL := fmt.Sprintf("https://releases.hashicorp.com/terraform/%v/terraform_%v_%v_%v.zip", version, version, platform, arch)

		file.Download(
			URL,
			localFileName,
		)

		fmt.Printf("[%v][Install] Decompress archive\n", name)
		err := archive.Unzip(localFileName, destFolder)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Delete archive\n", name)
		err = file.Delete(localFileName)
		if err != nil {
			panic(err)
		}

		// Rename file to terraform_${version}
		err = os.Rename(
			fmt.Sprintf(
				"%v/terraform",
				destFolder,
			),
			fmt.Sprintf(
				"%v/terraform_%v",
				destFolder,
				version,
			),
		)
		if err != nil {
			panic(err)
		}

		// Link terraform_${version} to terraform
		err = os.Symlink(
			fmt.Sprintf(
				"%v/terraform_%v",
				destFolder,
				version,
			),
			fmt.Sprintf(
				"%v/terraform",
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
