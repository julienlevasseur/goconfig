package vscodium

import (
	"fmt"

	"github.com/julienlevasseur/GoCfgMgr/pkg/system"
)

// Install VSCodium
func Install(version, arch, platform string, notIf ...bool) {
	if !notIf[0] {
		fmt.Println("[VSCodium][Install] Download archive")

		distribution, err := system.Distribution()
		if err != nil {
			panic(err)
		}

		fmt.Println(distribution)

		//localFileName := fmt.Sprintf("traefik_%s_%s_%s.tar.gz", arch, platform, version)
		//URL := fmt.Sprintf("https://github.com/traefik/traefik/releases/download/%s/traefik_%s_%s_%s.tar.gz", version, version, platform, arch)

		//file.Download(
		//	URL,
		//	localFileName,
		//)
	} else {
		fmt.Println("[VSCodium][Install] Ignore Install VSCodium due to NotIf")
	}
}
