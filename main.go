package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		dir    = flag.String("dir", ".", "target file path")
		bucket = flag.String("bucket", "hoge", "upload gcs bucket")
	)
	flag.Parse()
	fmt.Println(*dir)
	fmt.Println(*bucket)

	l, err := ListFiles(*dir)
	if err != nil {
		fmt.Printf("failed list files... %+v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()

	gcs, err := NewStorageService(ctx)
	if err != nil {
		fmt.Printf("failed NewStorageService... %+v\n", err)
		os.Exit(1)
	}

	c := NewConverter(gcs, *dir)
	if err := c.Run(ctx, *bucket, l, []string{}); err != nil {
		fmt.Printf("failed Converter.Run... %+v\n", err)
		os.Exit(1)
	}
}
