package cni

import (
	"fmt"
	"os"

	"github.comjulienlevasseur/GoCfgMgr/pkg/archive"
	"github.comjulienlevasseur/GoCfgMgr/pkg/file"
)

// Install CNI plugins
// version: CNI plugins version (v0.0.0)
// arch: host system architecture (amd64)
func Install(version, arch string) {
	localFileName := fmt.Sprintf("cni-plugins-linux-%s-%s.tgz", arch, version)
	URL := fmt.Sprintf("https://github.com/containernetworking/plugins/releases/download/%s/cni-plugins-linux-%s-%s.tgz", version, arch, version)

	file.Download(
		URL,
		localFileName,
	)

	tar, err := os.Open(localFileName)
	if err != nil {
		panic(err)
	}
	archive.Untar("/opt/cni/bin", tar)

	file.Delete(localFileName)
}
