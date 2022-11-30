package systemd

import (
	"fmt"

	"github.com/julienlevasseur/GoCfgMgr/pkg/file"
)

type Service struct {
	ServiceType string // user, system
	Name        string
	Description string
	Requires    string
	After       string
	Environment string
	Restart     string
	ExecStart   string
	ExecReload  string
	ExecStop    string
	KillSignal  string
	WantedBy    string
}

func (s Service) New() {
	content := `[Unit]
{{if .Description}}
Description={{ .Description}}
{{end}}
{{if .Requires}}
Requires={{ .Requires}}
{{end}}
{{if .After}}
After={{ .After}}
{{end}}
		
[Service]
{{if .Environment}}
Environment={{ .Environment}}
{{end}}
{{if .Restart}}
Restart={{ .Restart}}
{{end}}

{{if .ExecStart}}
ExecStart={{ .ExecStart}}
{{end}}
{{if .ExecReload}}
ExecReload={{ .ExecReload}}
{{end}}
{{if .ExecStop}}
ExecStop={{ .ExecStop}}
{{end}}
{{if .KillSignal}}
KillSignal={{ .KillSignal}}
{{end}}

[Install]
{{if .WantedBy}}
WantedBy={{ .WantedBy}}
{{end}}
`

	err := file.Template(
		fmt.Sprintf("/etc/systemd/%v/%v.service", .ServiceType, .Name),
		content,
		s,
	)
	if err != nil {
		panic(err)
	}
}
