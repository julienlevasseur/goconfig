package apt

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/command"
	apt "github.com/sigmonsays/go-apt-client"
)

type IPackage struct {
	Name string
}

type Package apt.Package

func Update() error {
	fmt.Printf("[apt] update package index files\n")
	out, err := apt.CheckForUpdates()
	if err != nil {
		return err
	}
	fmt.Printf("\n%v\n", string(out))

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
		fmt.Printf("[%v][Install] installing package\n", p.Name)
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
