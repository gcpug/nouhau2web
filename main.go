package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var (
		dir = flag.String("dir", "default", "string flag")
	)
	flag.Parse()
	fmt.Println(*dir)

	if *dir == "" {
		*dir = "."
	}

	fi, err := ioutil.ReadDir(*dir)
	if err != nil {
		fmt.Printf("failed read dir : %+v", err)
	}
	for _, f := range fi {
		fmt.Println(f.Name())
	}
}
