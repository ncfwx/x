package fileutil

import(
	"testing"
)

func TestIsDir(t *testing.T) {
	path := "../fileutil"
	result := IsDir(path)
	if !result {
		t.Errorf("test exists dir failed.")
	}

	path = "/"
	result = IsDir(path)
	if !result {
		t.Errorf("test exists dir failed.")
	}

	path = "./xxxx"
	result = IsDir(path)
	if result {
		t.Errorf("test no exists dir failed.")
	}

}
