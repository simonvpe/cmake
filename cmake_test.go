package cmake

import (
	"io/ioutil"
	"testing"
	"os"
	"fmt"
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

func TestGenerateMain(t *testing.T) {
	file := genTemporaryFile(t);
	defer os.Remove(file)

	ctx := CMakeContext {
		ProjectName:    "test-project",
		MinimumVersion: "3.8",
		TestSuite:      "catch",
	}
	
	err := Generate(file, "CMakeLists.tmpl", &ctx)
	if err != nil {
		t.Errorf("Error %q", err)
	}
	
	{
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			t.Error("Failed to read bytes")
		}
		fmt.Println(string(bytes))
	}
}

