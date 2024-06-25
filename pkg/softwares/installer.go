package softwares

type Installer interface {
	// Install(name, version, arch, platform string, notIf ...bool)
	InstallFromArchiveURL(version, arch, platform string, notIf bool)
	InstallFromURLBinaryURL(BinaryURL, destFolder string, notIf ...bool)
	NotIf(condition string) bool
}
