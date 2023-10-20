package cni

import (
	"fmt"
	"os"

	"github.com/julienlevasseur/goconfig/pkg/archive"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

// Install CNI plugins
// version: CNI plugins version (v0.0.0)
// arch: host system architecture (amd64)
func Install(version, arch string) {
	fmt.Println("[CNI][Install] Download archive")
	localFileName := fmt.Sprintf("cni-plugins-linux-%s-%s.tgz", arch, version)
	URL := fmt.Sprintf("https://github.com/containernetworking/plugins/releases/download/%s/cni-plugins-linux-%s-%s.tgz", version, arch, version)

	file.Download(
		URL,
		localFileName,
	)

	fmt.Println("[CNI][Install] Uncompress archive")
	tar, err := os.Open(localFileName)
	if err != nil {
		panic(err)
	}
	archive.Untar(tar, "/opt/cni/bin")

	fmt.Println("[CNI][Install] Delete archive")
	file.Delete(localFileName)

	fmt.Println("[CNI][Install] Complete")
}
