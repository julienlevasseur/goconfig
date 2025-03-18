package user

import (
	"os"
	"os/user"
	"strconv"
)

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

func (u *User) UID() (string, error) {
	us, err := user.Lookup(u.Username)
	if err != nil {
		return "", err
	}

	return us.Uid, nil
}

func (u *User) GID() (string, error) {
	us, err := user.Lookup(u.Username)
	if err != nil {
		return "", err
	}

	return us.Gid, nil
}

func (u *User) ChownToUser(path string) error {
	strUID, err := u.UID()
	if err != nil {
		return err
	}

	uid, err := strconv.Atoi(strUID)
	if err != nil {
		return err
	}

	strGID, err := u.GID()
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(strGID)
	if err != nil {
		return err
	}

	err = os.Chown(path, uid, gid)
	if err != nil {
		return err
	}

	return nil
}

func Username() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.Username, nil
}
