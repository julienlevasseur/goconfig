package softwares

type Software struct {
	Name    string
	Version string
}

var s Software

func (s *Software) Install() error {
	return s.Install()
}
