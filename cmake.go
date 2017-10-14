package cmake

import "os"

type CMakeContext struct {
	File          *os.File
	ProjectName    string
	MinimumVersion string
}

func Generate(filename, proj, ver string) (*CMakeContext, error) {
	f, err := os.Create(filename)
	if err != nil || f == nil {
		return nil, err
	}
	ctx := CMakeContext {
		File: f,
		ProjectName: proj,
		MinimumVersion: ver,
	}
	return &ctx, nil
}
