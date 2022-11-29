package consul

import (
	"fmt"
	"path/filepath"

	"github.com/julienlevasseur/GoCfgMgr/pkg/archive"
	"github.com/julienlevasseur/GoCfgMgr/pkg/file"
)

// Install Consul
// version: Consul version (v0.0.0)
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(version, arch, platform string, notIf ...bool) {
	if !notIf[0] {
		fmt.Println("[Consul][Install] Download archive")
		localFileName := fmt.Sprintf("consul_%s_%s_%s.zip", version, platform, arch)
		URL := fmt.Sprintf("https://releases.hashicorp.com/consul/%s/consul_%s_%s_%s.zip", version, version, platform, arch)

		file.Download(
			URL,
			localFileName,
		)

		fmt.Println("[Consul][Install] Uncompress archive")
		path, err := filepath.Abs(localFileName)
		if err != nil {
			panic(err)
		}
		archive.Unzip(path, "/usr/local/bin")

		fmt.Println("[Consul][Install] Delete archive")
		//file.Delete(localFileName)

		fmt.Println("[Consul][Install] Complete")
	} else {
		fmt.Println("[Consul][Install] Ignore InstallConsul due to NotIf")
	}
}
