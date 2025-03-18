package teams

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/dpkg"
	"github.com/julienlevasseur/goconfig/pkg/file"
	"github.com/julienlevasseur/goconfig/pkg/notif"
	"github.com/julienlevasseur/goconfig/pkg/softwares"
	"github.com/julienlevasseur/goconfig/pkg/user"
)

type Teams softwares.Software

func (t *Teams) Install() error {
	fmt.Printf("\n[%v][Install][Package]", t.Name)
	teams := dpkg.Dpkg{
		Name: "teams",
		ArchiveURL: fmt.Sprintf(
			"https://github.com/IsmaelMartinez/teams-for-linux/releases/download/v%v/teams-for-linux_%v_amd64.deb",
			t.Version,
			t.Version,
		),
		NotIf: notif.FileExists("/usr/bin/teams-for-linux"),
	}
	err := teams.Install()
	if err != nil {
		return err
	}

	fmt.Printf("\n[%v][Install][Menu Entry]", t.Name)

	var homePath string
	if t.Options != nil && t.Options.Username != nil {
		u := user.User{Username: "julien"}
		homePath, err = u.HomeDir()
	} else {
		u := user.User{}
		homePath, err = u.HomeDir()
	}

	err = file.Append(
		homePath+"/.local/share/applications/teams-for-linux.desktop",
		`[Desktop Entry]
Type=Application
Icon=/usr/share/icons/hicolor/1024x1024/apps/teams-for-linux.png
Name=Teams for Linux
Comment=Unofficial Microsoft Teams client for Linux
Categories=Office
Exec=/opt/teams-for-linux/teams-for-linux --no-sandbox %U
#Path=<set working directory if necessary>
StartupNotify=true
Terminal=false
		`,
	)

	return err
}
