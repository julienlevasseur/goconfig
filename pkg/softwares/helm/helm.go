package helm

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/apt"
	"github.com/julienlevasseur/goconfig/pkg/command"
)

const name = "helm"

// Install helm
func Install(notIf ...bool) {
	if !notIf[0] {
		fmt.Printf("\n[%v][Install]\n", name)
		command.Exec(
			"bash",
			[]string{
				"-c",
				"curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null",
			},
		)

		apt.Update()
		apt.Packages([]string{"apt-transport-https"})

		command.Exec(
			"bash",
			[]string{
				"-c",
				"echo \"deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main\" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list",
			},
		)

		apt.Update()
		apt.Packages([]string{"helm"})
	} else {
		fmt.Printf("[%v][Install] Ignore Installation due to NotIf\n", name)
	}
}
