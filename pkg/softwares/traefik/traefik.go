package traefik

import (
	"fmt"
	"os"

	"github.com/julienlevasseur/GoCfgMgr/pkg/archive"
	"github.com/julienlevasseur/GoCfgMgr/pkg/file"
)

// Install Traefik
// version: Traefik version (v0.0.0)
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(version, arch, platform string, notIf ...bool) {
	if !notIf[0] {
		fmt.Println("[Traefik][Install] Download archive")
		localFileName := fmt.Sprintf("traefik_%s_%s_%s.tar.gz", arch, platform, version)
		URL := fmt.Sprintf("https://github.com/traefik/traefik/releases/download/%s/traefik_%s_%s_%s.tar.gz", version, version, platform, arch)

		file.Download(
			URL,
			localFileName,
		)

		fmt.Println("[Traefik][Install] Uncompress archive")
		tar, err := os.Open(localFileName)
		if err != nil {
			panic(err)
		}
		archive.Untar("/usr/local/bin", tar)

		fmt.Println("[Traefik][Install] Delete archive")
		file.Delete(localFileName)

		fmt.Println("[Traefik][Install] Complete")
	} else {
		fmt.Println("[Traefik][Install] Ignore InstallTraefik due to NotIf")
	}
}
