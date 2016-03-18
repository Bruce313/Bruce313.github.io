package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const md_DIR = "./md"

func main() {
	files, err := ioutil.ReadDir(md_DIR)
	if err != nil {
		log.Fatal("read dir:", err)
	}
	for _, f := range files {
		fmt.Printf("read file %s\n", f.Name())
	}
}
