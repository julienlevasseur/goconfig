package user

import "os/user"

func Username() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
