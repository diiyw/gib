package orm

import (
	"testing"
)

func TestOpen(t *testing.T) {

	orm, err := Open(
		Auth("root", ""),
	)
	if err != nil {
		t.Fatal(err)
	}
	_ = orm
}

type Test struct {
	v interface{}
}