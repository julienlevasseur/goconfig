package git

import (
	"os"
	"text/template"

	"github.com/julienlevasseur/goconfig/pkg/user"
)

func GitIgnoreUser(ignore []string, user user.User) error {
	var tmplFile = "gitignore.tmpl"

	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	homeDir, err := user.HomeDir()
	if err != nil {
		return err
	}

	f, err := os.Create(homeDir + "/.gitignore")
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, ignore)
	if err != nil {
		return err
	}

	return nil
}
