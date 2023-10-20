package nomad

import (
	"fmt"
	"path/filepath"

	"github.com/julienlevasseur/GoCfgMgr/pkg/archive"
	"github.com/julienlevasseur/GoCfgMgr/pkg/file"
)

type Nomad struct {
	Version  string
	Arch     string
	Platform string
}

// Install Nomad
// version: Nomad version (v0.0.0)
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(nomad Nomad, notIf ...bool) {
	if !notIf[0] {
		fmt.Println("[Nomad][Install] Download archive")
		localFileName := fmt.Sprintf("nomad_%s_%s_%s.zip", nomad.Version, nomad.Platform, nomad.Platform)
		URL := fmt.Sprintf("https://releases.hashicorp.com/nomad/%s/nomad_%s_%s_%s.zip", nomad.Version, nomad.Version, nomad.Version, nomad.Version)

		file.Download(
			URL,
			localFileName,
		)

		fmt.Println("[Nomad][Install] Uncompress archive")
		path, err := filepath.Abs(localFileName)
		if err != nil {
			panic(err)
		}
		archive.Unzip(path, "/usr/local/bin")

		fmt.Println("[Nomad][Install] Delete archive")
		//file.Delete(localFileName)

		fmt.Println("[Nomad][Install] Complete")
	} else {
		fmt.Println("[Nomad][Install] Ignore InstallNomad due to NotIf")
	}
}
