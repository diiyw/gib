package gos

import (
	"fmt"
	"testing"
)

func TestGetDirTree(t *testing.T) {
	tree, err := Dirs("../")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tree)
}
