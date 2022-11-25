package file

import (
	"html/template"
	"io"
	"net/http"
	"os"
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

func New(path string) error {
	_, err := os.Create(path)
	return err
}

func Template(path, content string, vars any) error {
	t, err := template.New(path).Parse(content)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return t.Execute(file, vars)
}
