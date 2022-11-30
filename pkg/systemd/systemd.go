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

func (s Service) New(name, serviceType string, service Service) {
	content := `[Unit]
{{if s.Description}}
Description={{ s.Description}}
{{end}}
{{if s.Requires}}
Requires={{ s.Requires}}
{{end}}
{{if s.After}}
After={{ s.After}}
{{end}}
		
[Service]
{{if s.Environment}}
Environment={{ s.Environment}}
{{end}}
{{if s.Restart}}
Restart={{ s.Restart}}
{{end}}

{{if s.ExecStart}}
ExecStart={{ s.ExecStart}}
{{end}}
{{if s.ExecReload}}
ExecReload={{ s.ExecReload}}
{{end}}
{{if s.ExecStop}}
ExecStop={{ s.ExecStop}}
{{end}}
{{if s.KillSignal}}
KillSignal={{ s.KillSignal}}
{{end}}

[Install]
{{if s.WantedBy}}
WantedBy={{ s.WantedBy}}
{{end}}
`

	file.Template(
		fmt.Sprintf("/etc/systemd/%v/%v.service", s.serviceType, s.name),
		content,
		service,
	)
}
