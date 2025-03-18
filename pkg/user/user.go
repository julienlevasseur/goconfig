package user

import "os/user"

type User struct {
	Username string
}

var u User

func (u *User) HomeDir() (string, error) {
	if u.Username != "" {
		us, err := user.Lookup(u.Username)
		if err != nil {
			return "", err
		}

		return us.HomeDir, nil
	} else {
		us, err := user.Current()
		if err != nil {
			return "", err
		}
		return us.HomeDir, nil
	}

}

func Username() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.Username, nil
}
