package system

import "os/exec"

func Distribution() (string, error) {
	out, err := exec.Command(
		"grep",
		"DISTRIB_ID",
		"/etc/lsb-release|cut",
		"-d",
		"'='",
		"-f",
		"2",
	).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
