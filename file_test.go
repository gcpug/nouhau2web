package main

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestListFiles(t *testing.T) {
	l, err := ListFiles("/Users/sinmetal/go/src/github.com/sinmetal/nouhau")
	if err != nil {
		t.Fatal(err)
	}
	pp.Println(l)
}
