package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var (
		dir = flag.String("dir", ".", "target file path")
	)
	flag.Parse()
	fmt.Println(*dir)

	fi, err := ioutil.ReadDir(*dir)
	if err != nil {
		fmt.Printf("failed read dir : %+v", err)
	}
	for _, f := range fi {
		fmt.Println(f.Name())
	}
}
