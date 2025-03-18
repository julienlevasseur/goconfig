package apt

import (
	"fmt"
	"os"
	"time"

	"github.com/julienlevasseur/goconfig/pkg/command"
	apt "github.com/sigmonsays/go-apt-client"
)

type IPackage struct {
	Name string
}

type Package apt.Package

func Update() error {
	ok, err := lastUpdateInLast24h()
	if err != nil {
		return err
	}

	if !ok {
		fmt.Printf("[apt] Update package index files\n")
		out, err := apt.CheckForUpdates()
		if err != nil {
			return err
		}
		fmt.Printf("\n%v\n", string(out))
	} else {
		fmt.Printf("[apt] Ingoring Update due to last update < 24h\n")
	}

	return err
}

// Packages takes a list of package names (slice of string) to install
func Packages(names []string) error {
	for _, name := range names {
		fmt.Printf("[apt][%v][install] Installing ...\n", name)
		pkg := IPackage{
			Name: name,
		}

		// Don't check for error here, since a unsuccessful grep will return an error status 1:
		installed, _ := pkg.Installed()

		err := pkg.Install(installed)
		if err != nil {
			return err
		}
		fmt.Printf("[apt][%v][install] Complete\n", name)
	}

	return nil
}

func (p IPackage) Install(notIf *bool) error {
	if notIf != nil && !*notIf {
		fmt.Printf("[%v][Install] Installing package\n", p.Name)
		pkg := apt.Package{
			Name: p.Name,
		}

		out, err := apt.Install(&pkg)
		if err != nil {
			return err
		}
		fmt.Printf("\n%v\n", string(out))

		fmt.Printf("[%v][Install] Complete\n", p.Name)
		return err
	} else {
		fmt.Printf("[%v][Install] Ignore installation due to NotIf\n", p.Name)
	}

	return nil
}

func (p IPackage) Installed() (*bool, error) {
	output, err := command.ExecWithOutput(
		"bash",
		[]string{
			"-c",
			fmt.Sprintf("dpkg --list|grep %v", p.Name),
		},
	)
	if output != "" {
		t := true
		return &t, err
	}

	f := false
	return &f, err
}

func getLastUpdate() (time.Time, error) {
	fileInfo, err := os.Stat("/var/lib/apt/lists")
	if err != nil {
		return time.Time{}, err
	}

	return fileInfo.ModTime(), nil
}

func lastUpdateInLast24h() (bool, error) {
	now := time.Now()

	lastUpdate, err := getLastUpdate()
	if err != nil {
		return false, err
	}

	yesterday := now.Add(-24 * time.Hour)
	if lastUpdate.After(yesterday) {
		return true, nil
	}

	return false, nil
}
