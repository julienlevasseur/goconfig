package softwares

type InstallOptions struct {
	Username *string
}

type Software struct {
	Name    string
	Version string
	Options *InstallOptions
}

var s Software

func (s *Software) Install() error {
	return s.Install()
}
