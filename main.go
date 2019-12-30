package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	var (
		dir = flag.String("dir", ".", "target file path")
	)
	flag.Parse()
	fmt.Println(*dir)

	l, err := ListFiles(*dir)
	if err != nil {
		fmt.Printf("failed list files... %+v\n", err)
		os.Exit(1)
	}
	_, err = pp.Println(l)
	if err != nil {
		fmt.Printf("failed k0kubun pp... %+v\n", err)
		os.Exit(1)
	}

	//ctx := context.Background()
	//
	//gcs, err := NewStorageService(ctx)
	//if err != nil {
	//	fmt.Printf("failed NewStorageService... %+v\n", err)
	//	os.Exit(1)
	//}

	//	w := gcs.NewWriter(ctx, "gcpug-nouhau-dev", "")
	//w.ObjectAttrs.ContentType = "text/csv; charset=utf-8"
	//w.ObjectAttrs.ContentDisposition = fmt.Sprintf(`attachment;filename="%v-%v.csv"`, runRequest.ID, stmt.ID)

	// io.Copy().,0
}
