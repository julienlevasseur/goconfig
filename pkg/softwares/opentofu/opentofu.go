package opentofu

import (
	"fmt"
	"log"

	"github.com/julienlevasseur/goconfig/pkg/apt"
	"github.com/julienlevasseur/goconfig/pkg/command"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

const name = "opentofu"

// Install helm
func Install(notIf ...bool) {
	if !notIf[0] {
		fmt.Printf("\n[%v][Install]\n", name)

		apt.Update(false)
		apt.Packages([]string{
			"apt-transport-https",
			"ca-certificates",
			"curl",
			"gnupg",
		})

		args := []string{
			"-c",
			"install -m 0755 -d /etc/apt/keyrings",
		}
		command.Exec(
			"bash",
			&args,
		)

		args = []string{
			"-c",
			"curl -fsSL https://get.opentofu.org/opentofu.gpg | sudo tee /etc/apt/keyrings/opentofu.gpg >/dev/null",
		}
		command.Exec(
			"bash",
			&args,
		)

		args = []string{
			"-c",
			"curl -fsSL https://packages.opentofu.org/opentofu/tofu/gpgkey | sudo gpg --no-tty --batch --dearmor -o /etc/apt/keyrings/opentofu-repo.gpg >/dev/null",
		}
		command.Exec(
			"bash",
			&args,
		)

		args = []string{
			"-c",
			"chmod a+r /etc/apt/keyrings/opentofu.gpg /etc/apt/keyrings/opentofu-repo.gpg",
		}
		command.Exec(
			"bash",
			&args,
		)

		file.Content(
			"/etc/apt/sources.list.d/opentofu.list",
			`deb [signed-by=/etc/apt/keyrings/opentofu.gpg,/etc/apt/keyrings/opentofu-repo.gpg] https://packages.opentofu.org/opentofu/tofu/any/ any main
deb-src [signed-by=/etc/apt/keyrings/opentofu.gpg,/etc/apt/keyrings/opentofu-repo.gpg] https://packages.opentofu.org/opentofu/tofu/any/ any main`,
			file.Exists("/etc/apt/sources.list.d/opentofu.list"),
			nil,
		)

		apt.Update(true)
		tofuPkg := apt.IPackage{
			Name: "tofu",
		}
		err := tofuPkg.Install(
			file.Exists("/usr/bin/tofu"),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
