package prusaslicer

import (
	"fmt"
	"strings"

	"github.com/julienlevasseur/goconfig/pkg/archive"
	"github.com/julienlevasseur/goconfig/pkg/command"
	"github.com/julienlevasseur/goconfig/pkg/file"
)

type PrusaSlicerInstaller struct{}

func (psi *PrusaSlicerInstaller) Install(version, platform, arch string, notIf bool) {
	if notIf {
		fmt.Println("[PrusaSlicer][Install] Ignore Install due to NotIf")
	} else {
		fmt.Println("[PrusaSlicer][Install] Download archive")
		v := strings.Replace(version, ".", "_", -1)
		URL := fmt.Sprintf("https://cdn.prusa3d.com/downloads/drivers/prusa3d_%v_%v.zip", platform, v)
		localFileName := fmt.Sprintf("prusa3d_%v_%v.zip", platform, v)

		err := file.Download(
			URL,
			localFileName,
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("[PrusaSlicer][Install] Decompress archive")
		err = archive.Unzip(localFileName, "/tmp")
		if err != nil {
			panic(err)
		}

		//err = file.Move(
		//	fmt.Sprintf(
		//		"/tmp/PrusaSlicer-%v+%v-%v-GTK3-*.AppImage",
		//		version,
		//		platform,
		//		arch,
		//	),
		//	fmt.Sprintf(
		//		"/opt/PrusaSlicer-%v+%v-%v-GTK3.AppImage",
		//		version,
		//		platform,
		//		arch,
		//	),
		//)
		command.Exec(
			fmt.Sprintf(
				"mv /tmp/PrusaSlicer-%v+%v-%v-GTK3-*.AppImage /opt/PrusaSlicer-%v+%v-%v-GTK3.AppImage",
				version,
				platform,
				arch,
				version,
				platform,
				arch,
			),
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("[PrusaSlicer][Install] Delete archive")
		err = file.Delete(localFileName)
		if err != nil {
			panic(err)
		}

		fmt.Println("[PrusaSlicer][Install] Complete")
	}
}
