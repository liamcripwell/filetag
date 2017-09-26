package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Record - represents a file record
type Record struct {
	FileName string   `json:"filename"`
	Tags     []string `json:"tags"`
}

func main() {
	// get command argument
	command := os.Args[1]

	// read existing records from disk
	records := getRecordsFromFile("data.json")

	// perform chosen command
	// TODO thing
	switch command {
	case "tag":
		addContent(records)
	case "untag":
		removeContent(records)
	}

}

func removeContent(records []Record) {
	for i, record := range records {
		if record.FileName == os.Args[2] {
			records = append(records[:i], records[i+1:]...)
		}
	}

	// write updated data to file
	writeRecordsToFile(records, "data.json")
}

// Adds new content to the data store
func addContent(records []Record) {
	// check whether a record already exists for file
	existing := -1
	for i, record := range records {
		if record.FileName == os.Args[2] {
			existing = i
		}
	}

	// update records with new content
	if existing < 0 {
		newFile := Record{os.Args[2], os.Args[3:]}
		records = append(records, newFile)
	} else {
		for _, tag := range os.Args[3:] {
			if !stringInSlice(tag, records[existing].Tags) {
				records[existing].Tags =
					append(records[existing].Tags, tag)
			}
		}
	}

	// write updated data to file
	writeRecordsToFile(records, "data.json")
}

// Compiles a Record slice from a given JSON file
func getRecordsFromFile(file string) []Record {
	// read in file
	b, _ := ioutil.ReadFile(file)
	jsonContent := string(b)

	// parse the JSON
	var objs []Record
	json.Unmarshal([]byte(jsonContent), &objs)

	return objs
}

// Writes a Record slice to a given JSON file
func writeRecordsToFile(records []Record, file string) {
	// open file to write data to
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Can't open data file...")
		return
	}
	defer f.Close()

	// marshal files to json
	data, err := json.MarshalIndent(records, "", "	")
	if err != nil {
		fmt.Println("Can't create json...")
	}

	fmt.Println(string(data))

	// write data to file
	f.Write(data)
	f.Close()
}

// Checks whether a given string exists within a given slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
