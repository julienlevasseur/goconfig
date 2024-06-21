package softwares

type Installer interface {
	Install(name, version, arch, platform string, notIf ...bool)
}
