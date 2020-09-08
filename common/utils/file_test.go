package utils

import (
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	dir := "/Users/fante/devs/projects/testmover"
	files, err := GetAllFiles(dir)
	if err != nil {
		t.Errorf("Get Dir %s Error %v\n", dir, err)
	} else {
		t.Log(files)
	}
}
