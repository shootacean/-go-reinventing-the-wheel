package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	files, _ := ioutil.ReadDir(wd)
	for _, f := range files {
		fmt.Printf("%s\t", f.Name())
	}
	fmt.Println("")
}
