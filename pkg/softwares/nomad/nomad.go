package nomad

import (
	"fmt"
	"path/filepath"

	"github.com/julienlevasseur/goconfig/pkg/archive"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

const name = "nomad"

// Install Nomad
// version: Nomad version (v0.0.0)
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(version, arch, platform string, notIf ...bool) {
	if !notIf[0] {
		fmt.Printf("\n[%v][Install] Download archive\n", name)
		localFileName := fmt.Sprintf("nomad_%s_%s_%s.zip", version, platform, arch)
		URL := fmt.Sprintf("https://releases.hashicorp.com/nomad/%s/nomad%s_%s_%s.zip", version, version, platform, arch)

		file.Download(
			URL,
			localFileName,
		)

		fmt.Printf("\n[%v][Install] Uncompress archive\n", name)
		path, err := filepath.Abs(localFileName)
		if err != nil {
			panic(err)
		}
		archive.Unzip(path, "/usr/local/bin")

		fmt.Printf("\n[%v][Install] Delete archive\n", name)

		fmt.Printf("\n[%v][Install] Complete\n", name)
	} else {
		fmt.Printf("\n[%v][Install] Ignore InstallConsul due to NotIf\n", name)
	}
}
