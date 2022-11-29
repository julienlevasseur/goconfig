package file

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Append(path, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func Delete(path string) error {
	return os.Remove(path)
}

func Download(URL, fileDest string) error {
	out, err := os.Create(fileDest)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	} else {
		return true, nil
	}
}

func LineIsPresent(path, match string) (bool, error) {
	fileExists, err := Exists(path)
	if err != nil {
		return false, err
	}
	if fileExists {
		f, err := ioutil.ReadFile(path)
		if err != nil {
			return false, err
		}

		lines := strings.Split(string(f), "\n")

		for _, line := range lines {
			if strings.Contains(line, match) {
				return true, nil
			}
		}
	}

	return false, nil
}

func New(path string) error {
	_, err := os.Create(path)
	return err
}

func ReplaceLine(path, match, content string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(f), "\n")

	for i, line := range lines {
		if strings.Contains(line, match) {
			lines[i] = content
		}
	}

	out := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(out), 0644)
	if err != nil {
		return err
	}

	return nil
}

func Template(target, content string, vars any) error {
	fmt.Println("[File][Template] create template ", target)
	t, err := template.New(target).Parse(content)
	if err != nil {
		return err
	}
	file, err := os.Create(target)
	if err != nil {
		return err
	}
	return t.Execute(file, vars)
}
