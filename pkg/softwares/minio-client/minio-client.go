package minioClient

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/dpkg"
	"github.com/julienlevasseur/goconfig/pkg/notif"
)

const name = "minio-client"

// Install Minio Client
// version: Minio Client version (v0.0.0)
// arch: host system architecture (amd64)
// platform: host system platform (linux)
func Install(version, arch, platform string, notIf ...bool) error {
	if !notIf[0] {
		fmt.Printf("\n[%v][Install] Install package\n", name)
		mc := dpkg.Dpkg{
			Name: "mc",
			ArchiveURL: fmt.Sprintf(
				"https://dl.min.io/aistor/minio/release/%v-%v/minio_%v_%v.deb",
				platform,
				arch,
				version,
				arch,
			),
			NotIf: notif.FileExists("/usr/bin/codium"),
		}
		err := mc.Install()
		if err != nil {
			return err
		}

		fmt.Printf("\n[%v][Install] Delete archive\n", name)

		fmt.Printf("\n[%v][Install] Complete\n", name)
	} else {
		// notif.IgnoreDueToNotIf(name, "Install")
		fmt.Printf("\n[%v][Install] Ignored due to NotIf\n", name)

	}

	return nil
}
