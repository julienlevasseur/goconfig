package file

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

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
	fmt.Println(vars)
	return t.Execute(file, vars)
}
