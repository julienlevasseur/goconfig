package docker

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/command"
)

const name = "docker"

// Install Dcoker
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(arch, platform string, notIf ...bool) {
	if !notIf[0] {
		fmt.Printf("[%v][Install] Add Docker's GPGP key\n", name)

		args := []string{
			"-m",
			"0755",
			"-d",
			"/etc/apt/keyrings",
		}
		err := command.Exec(
			"install",
			&args,
		)
		if err != nil {
			fmt.Printf("[Error]: %q\n", err)
		}

		args = []string{
			"-fsSL",
			"https://download.docker.com/linux/ubuntu/gpg",
			"-o",
			"/etc/apt/keyrings/docker.asc",
		}
		err = command.Exec(
			"curl",
			&args,
		)
		if err != nil {
			fmt.Printf("[Error]: %q\n", err)
		}

		args = []string{
			"a+r",
			"/etc/apt/keyrings/docker.asc",
		}
		err = command.Exec(
			"chmod",
			&args,
		)
		if err != nil {
			fmt.Printf("[Error]: %q\n", err)
		}

		fmt.Printf("\n[%v][Install] Add Docker's repository to APT sources\n", name)
		//echo \
		//  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
		//  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
		//  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

		// 		err = file.Append(
		// 			"/etc/apt/sources.list.d/docker.list",
		// 			`
		// deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
		//   $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable"
		// `,
		// 		)
		// 		if err != nil {
		// 			fmt.Printf("\n[Error]: %q\n", err)
		// 		}

		fmt.Printf("\n[%v][Install] Complete\n", name)
	} else {
		fmt.Printf("\n[%v][Install] Ignore installation due to NotIf\n", name)
	}
}
