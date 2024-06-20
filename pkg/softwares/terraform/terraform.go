package terraform

import (
	"fmt"
	"os"

	"github.com/julienlevasseur/goconfig/pkg/archive"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

const name = "Terraform"

func Install(version, arch, platform string, notIf ...bool) {
	if notIf[0] {
		fmt.Printf("[%v][Install] Ignore Install due to NotIf", name)
	} else {
		fmt.Printf("[%v][Install] Download archive", name)
		localFileName := fmt.Sprintf("terraform_%v_%v_%v.zip", version, version, platform, arch)
		URL := fmt.Sprintf("https://releases.hashicorp.com/terraform/%v/terraform_%v_%v_%v.zip", version, version, platform, arch)

		file.Download(
			URL,
			localFileName,
		)

		// fmt.Printf("[%v][Install] Decompress archive", name)
		// zip, err := os.Open(localFileName)
		// if err != nil {
		// 	panic(err)
		// }
		err := archive.Unzip(localFileName, "/usr/local/bin")
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Delete archive", name)
		err = file.Delete(localFileName)
		if err != nil {
			panic(err)
		}

		// Rename file to terraform_${version}
		err = os.Rename("/usr/local/bin/terraform", fmt.Sprintf("/usr/local/bin/terraform_%v", version))
		if err != nil {
			panic(err)
		}

		// Link terraform_${version} to terraform
		err = os.Symlink(fmt.Sprintf("/usr/local/bin/terraform_%v", version), "/usr/local/bin/terraform")
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Complete", name)
	}
}
