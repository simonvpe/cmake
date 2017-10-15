package cmake

import (
	"os"
	"text/template"
)

type CMakeContext struct {
	ProjectName    string
	MinimumVersion string
	TestSuite      string
}

func Generate(file, tmpl string, ctx *CMakeContext) (error) {
	t := template.Must(template.ParseFiles(tmpl))
	
	f, err := os.OpenFile(file, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	
	err = t.Execute(f, ctx)
	if err != nil {
		return err
	}
	
	return nil
}
