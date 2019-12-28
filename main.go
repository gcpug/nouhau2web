package main

import (
	"flag"
	"fmt"

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
	}
	_, err = pp.Println(l)
	if err != nil {
		fmt.Printf("failed k0kubun pp... %+v\n", err)
	}
}
