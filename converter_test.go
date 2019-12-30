package main

import (
	"context"
	"testing"
)

func TestConverter_LocalToGCS(t *testing.T) {
	c := newConverter(t)

	ctx := context.Background()

	fn := "README.md"
	lfp := "/Users/sinmetal/go/src/github.com/gcpug/nouhau2web/README.md"

	if err := c.Process(ctx, lfp, "gcpug-nouhau-test", fn); err != nil {
		t.Fatal(err)
	}
}

func TestConverter_ContentType(t *testing.T) {
	c := newConverter(t)

	ext, _ := c.ContentType("hoge.MD")
	if e, g := ".md", ext; e != g {
		t.Errorf("Ext want %v got %v", e, g)
	}
}

func TestConverter_ObjectPath(t *testing.T) {
	cases := []struct {
		name string
		root string
		lfp  string
		want string
	}{
		{"fullPath", "/Users/sinmetal/hoge", "/Users/sinmetal/hoge/fuga.MD", "fuga.html"},
		{"currentPathMarkdown", ".", "./hoge/fuga.MD", "hoge/fuga.html"},
		{"currentPathText", ".", "./hoge/fuga.txt", "hoge/fuga.txt"},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := newConverterWithRoot(t, tt.root)
			got := c.ObjectPath(tt.lfp)
			if got != tt.want {
				t.Errorf("want %s but got %s", tt.want, got)
			}
		})
	}
}

func newConverter(t *testing.T) *Converter {
	ctx := context.Background()

	gcs, err := NewStorageService(ctx)
	if err != nil {
		t.Fatal(err)
	}

	return NewConverter(gcs, ".")
}

func newConverterWithRoot(t *testing.T, root string) *Converter {
	ctx := context.Background()

	gcs, err := NewStorageService(ctx)
	if err != nil {
		t.Fatal(err)
	}

	return NewConverter(gcs, root)
}
