package cmake

import (
	"io/ioutil"
	"testing"
	"os"
)

func genTemporaryFile(t *testing.T) string {
	f, err := ioutil.TempFile("","test-cmake-")
	if err != nil || f == nil {
		t.Error("Failed to create temporary file")
	}
	if f.Close() != nil {
		t.Error("Failed to close temporary file")
	}
	return f.Name()
}

func TestGenerate(t *testing.T) {
	proj := "test-project"
	mver := "3.8"
	file := genTemporaryFile(t);
	defer os.Remove(file)
	
	ctx, err := Generate(file, proj, mver)
	if err != nil {
		t.Error("Error %q", err)
	}
	if ctx == nil {
		t.Error("Context was nil")
	}
	if ctx.ProjectName != proj {
		t.Error("Bad ProjectName")
	}
	if ctx.MinimumVersion != mver {
		t.Error("Bad MinimumVersion")
	}	
}

