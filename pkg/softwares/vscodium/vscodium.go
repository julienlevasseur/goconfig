package vscodium

import (
	"fmt"
	"log"

	"github.com/julienlevasseur/goconfig/pkg/command"
	"github.com/julienlevasseur/goconfig/pkg/dpkg"
	"github.com/julienlevasseur/goconfig/pkg/notif"
	"github.com/julienlevasseur/goconfig/pkg/softwares"
	"github.com/julienlevasseur/goconfig/pkg/user"
)

type VSCodium softwares.Software

func (v *VSCodium) Install() error {
	log.Printf("\n[%v][Install][Package]", v.Name)

	codium := dpkg.Dpkg{
		Name: "codium",
		ArchiveURL: fmt.Sprintf(
			"https://github.com/VSCodium/vscodium/releases/download/%v/codium_%v_amd64.deb",
			v.Version,
			v.Version,
		),
		NotIf: notif.FileExists("/usr/bin/codium"),
	}
	err := codium.Install()
	if err != nil {
		return err
	}

	log.Printf("[%v][Install][Installation complete]\n", v.Name)

	return nil
}

func InstallExtension(name string, u *user.User) (string, error) {
	args := []string{
		"--install-extension",
		name,
	}

	output, err := command.ExecAs(
		u,
		"codium",
		&args,
	)
	if err != nil {
		log.Fatal(err)
	}

	return output, nil
}
