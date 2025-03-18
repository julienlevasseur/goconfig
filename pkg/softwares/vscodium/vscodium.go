package vscodium

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/dpkg"
	"github.com/julienlevasseur/goconfig/pkg/notif"
	"github.com/julienlevasseur/goconfig/pkg/softwares"
)

type VSCodium softwares.Software

func (v *VSCodium) Install() error {
	fmt.Printf("\n[%v][Install][Package]", v.Name)

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

	return nil
}
