package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fi, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("failed read dir : %+v", err)
	}
	for _, f := range fi {
		fmt.Println(f.Name())
	}
}
