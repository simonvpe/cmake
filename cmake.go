package cmake

import (
	"os"
	"text/template"
)

type CMakeContext struct {
	File          *os.File
	ProjectName    string
	MinimumVersion string
}

func GenerateMain(filename, proj, ver string) (*CMakeContext, error) {
	ctx := CMakeContext {
		File: nil,
		ProjectName: proj,
		MinimumVersion: ver,
	}
	if err := fromTemplate(ctx, filename, "CMakeLists.tmpl"); err != nil {
		return nil, err
	}
	return &ctx, nil
}

func fromTemplate(ctx CMakeContext, target, tmpl string) error {
	t := template.Must(template.ParseFiles(tmpl))
	
	f, err := os.OpenFile(target, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	
	err = t.Execute(f, ctx)
	if err != nil {
		return err
	}
	
	return nil
}
