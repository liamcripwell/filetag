package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// File ... to store tags
type File struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func main() {

	// open file to write data to
	f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Can't data open file...")
		return
	}
	defer f.Close()

	// create new File from cmdline args
	var files []File
	newFile := File{os.Args[1], []string{os.Args[2]}}
	files = append(files, newFile)

	// marshal files to json
	b, err := json.Marshal(files)
	if err != nil {
		fmt.Println("Can't create json...")
	}

	// write data to file
	f.Write(b)
	f.Close()

	fmt.Println("Done...")
}
