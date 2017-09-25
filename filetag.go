package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

// Record ... to store tags
type Record struct {
	FileName string   `json:"filename"`
	Tags     []string `json:"tags"`
}

func main() {

	// open file to write data to
	f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Can't data open file...")
		return
	}
	defer f.Close()

	b, err := ioutil.ReadFile("data.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	jsonContent := string(b)

	// Parse the JSON.
	var objs interface{}
	json.Unmarshal([]byte(jsonContent), &objs)

	// Ensure that it is an array of objects.
	objArr, ok := objs.([]interface{})
	if !ok {
		log.Fatal("expected an array of objects")
	}

	// Handle each object as a map[string]interface{}.
	for i, obj := range objArr {
		obj, ok := obj.(map[string]interface{})
		if !ok {
			log.Fatalf("expected type map[string]interface{}, got %s", reflect.TypeOf(objArr[i]))
		}
		fmt.Printf("i=%d, o=%T\n", i, obj) // Do something with the object...
		fmt.Println(obj["filename"])
	}

	// create new File from cmdline args
	var files []Record
	newFile := Record{os.Args[1], os.Args[2:]}
	files = append(files, newFile)

	// marshal files to json
	data, err := json.Marshal(files)
	if err != nil {
		fmt.Println("Can't create json...")
	}

	// write data to file
	f.Write(data)
	f.Close()

	fmt.Println("Done...")
}
