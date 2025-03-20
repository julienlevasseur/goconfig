package minikube

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/command"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

const name = "minikuke"

// Install Minikube
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(arch, platform string, notIf ...bool) {
	if !notIf[0] {
		fmt.Printf("\n[%v][Install] Download installation binary\n", name)
		localFileName := fmt.Sprintf("minikube-%v-%v", platform, arch)
		URL := fmt.Sprintf("https://storage.googleapis.com/minikube/releases/latest/minikube-%v-%v", platform, arch)

		err := file.Download(
			URL,
			localFileName,
		)
		if err != nil {
			fmt.Printf("[Error]: %q\n", err)
		}

		fmt.Printf("[%v][Install] Execute installation binary\n", name)
		err = command.Exec(
			"install",
			[]string{
				"minikube-linux-amd64",
				"/usr/local/bin/minikube",
			},
			nil,
		)
		if err != nil {
			fmt.Printf("[Error]: %q\n", err)
		}

		fmt.Printf("\n[%v][Install] Delete installation binary\n", name)
		err = file.Delete(localFileName)
		if err != nil {
			fmt.Printf("\n[Error]: %q\n", err)
		}

		fmt.Printf("[%v][Install] Complete\n", name)
	} else {
		fmt.Printf("[%v][Install] Ignore installation due to NotIf\n", name)
	}
}
