package cmake

import (
	"os"
	"text/template"
	"errors"
	"path"
	"fmt"
)

type CMakeContext struct {
	ProjectName    string
	MinimumVersion string
	TestSuite      string
	Language       string
}

func Generate(file, tmpl string, ctx *CMakeContext) (error) {
	name := path.Base(tmpl)
	t, err := template.New(name).Funcs(
	        template.FuncMap{
			"BadLanguage": func(lang string) (bool, error) {
				msg := fmt.Sprintf("Language %q not supported!", lang)
				return false, errors.New(msg)
			},
			"BadTestSuite": func(suite string) (bool, error) {
				msg := fmt.Sprintf("Test suite %q not supported!", suite)
				return false, errors.New(msg)
			},
		}).ParseFiles(tmpl)

	if err != nil || t == nil {
		return err
	}
	
	f, err := os.OpenFile(file, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil || f == nil {
		return err
	}

	err = t.Execute(f, ctx)
	if err != nil {
		return err
	}
	
	return nil
}
