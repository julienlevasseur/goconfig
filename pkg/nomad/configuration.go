package nomad

import (
	"html/template"
	"os"

	"github.com/hashicorp/nomad/command/agent"
)

type NomadConfig agent.Config

func ConfigFile(cfg NomadConfig, path string) error {
	tmpl, err := template.ParseFiles("templates/config.tmpl")
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	tmpl.Execute(file, cfg)

	return nil
}
