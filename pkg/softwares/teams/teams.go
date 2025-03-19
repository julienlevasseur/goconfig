package teams

import (
	"fmt"
	"log"

	"github.com/julienlevasseur/goconfig/pkg/dpkg"
	"github.com/julienlevasseur/goconfig/pkg/file"
	"github.com/julienlevasseur/goconfig/pkg/notif"
	"github.com/julienlevasseur/goconfig/pkg/softwares"
	"github.com/julienlevasseur/goconfig/pkg/user"
)

type Teams softwares.Software

func (t *Teams) Install() error {
	log.Printf("[%v][Install][Package]\n", t.Name)
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

	log.Printf("[%v][Install][Menu Entry]\n", t.Name)

	var homePath string
	if t.Options != nil && t.Options.Username != nil {
		u := user.User{Username: "julien"}
		homePath, err = u.HomeDir()
	} else {
		u := user.User{}
		homePath, err = u.HomeDir()
	}

	teamsDesktopFilePath := homePath + "/.local/share/applications/teams-for-linux.desktop"
	err = file.Append(
		teamsDesktopFilePath,
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
		notif.FileExists(teamsDesktopFilePath),
	)

	log.Printf("[%v][Install][Installation complete]\n", t.Name)

	return err
}
