package ssh

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/julienlevasseur/goconfig/pkg/fs"
	"github.com/julienlevasseur/goconfig/pkg/user"
)

// SSH config

type Host struct {
	Host         string
	Hostname     string `mapstructure:"Hostname,omitempty"`
	User         string `mapstructure:"User,omitempty"`
	IdentityFile string `mapstructure:"IdentityFile,omitempty"`
}

type Config struct {
	Hosts []Host `mapstructure:"Host"`
	User  string
}

// func NewConfig() error {
func NewConfig(cfg Config) error {
	log.Printf("[ssh][config] Render config template\n")
	tmpl, err := template.ParseFiles("pkg/ssh/templates/config.tmpl")
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/home/%s/.ssh/config", cfg.User)

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	log.Printf("[ssh][config] Write config template to user's %s ssh dir\n", cfg.User)
	tmpl.Execute(f, cfg)

	u := user.User{Username: cfg.User}

	log.Printf("[ssh][config] Ensure %s is owner of its ssh config\n", cfg.User)
	err = u.ChownToUser(path)
	if err != nil {
		return err
	}

	log.Printf("[ssh][config] Set ssh config file mode to 0400\n")
	fs.Chmod(path, 0400)

	log.Printf("[ssh][config] Complete\n")
	return nil
}
