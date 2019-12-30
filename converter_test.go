package main

import (
	"context"
	"testing"
)

func TestConverter_LocalToGCS(t *testing.T) {
	ctx := context.Background()

	gcs, err := NewStorageService(ctx)
	if err != nil {
		t.Fatal(err)
	}

	c := NewConverter(gcs)

	fn := "README.md"
	lfp := "/Users/sinmetal/go/src/github.com/gcpug/nouhau2web/README.md"

	if err := c.Process(ctx, lfp, "gcpug-nouhau-test", fn); err != nil {
		t.Fatal(err)
	}
}
