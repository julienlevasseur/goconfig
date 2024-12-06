package az_cli

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/apt"
)

const name = "az-cli"

// Install az-cli
// version: az-cli version (v0.0.0)
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(notIf ...bool) {
	if !notIf[0] {
		// fmt.Printf("[%v][Install] Download archive\n", name)
		// localFileName := fmt.Sprintf("consul_%s_%s_%s.zip", version, platform, arch)
		// URL := fmt.Sprintf("https://releases.hashicorp.com/consul/%s/consul_%s_%s_%s.zip", version, version, platform, arch)

		// err := file.Download(
		// 	URL,
		// 	localFileName,
		// )
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("[%v][Install] Uncompress archive\n", name)
		// path, err := filepath.Abs(localFileName)
		// if err != nil {
		// 	panic(err)
		// }
		// err = archive.Unzip(path, "/usr/local/bin")
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("[%v][Install] Delete archive\n", name)
		// //file.Delete(localFileName)

		err := apt.Update()
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%v][Install] Complete\n", name)
	} else {
		fmt.Printf("[%v][Install] Ignore InstallConsul due to NotIf\n", name)
	}
}
