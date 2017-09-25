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

	records := getRecordsFromFile("data.json")
	existing := -1

	// check whether a record already exists for file
	for i, record := range records {
		if record.FileName == os.Args[1] {
			existing = i
		}
	}

	// update records with new content
	if existing < 0 {
		newFile := Record{os.Args[1], os.Args[2:]}
		records = append(records, newFile)
	} else {
		for _, tag := range os.Args[2:] {
			if !stringInSlice(tag, records[existing].Tags) {
				records[existing].Tags =
					append(records[existing].Tags, tag)
			}
		}
	}

	// write updated data to file
	writeRecordsToFile(records, "data.json")
	fmt.Println("Done...")
}

func getRecordsFromFile(file string) []Record {
	// read in file
	b, _ := ioutil.ReadFile(file)
	jsonContent := string(b)

	// parse the JSON
	var objs []Record
	json.Unmarshal([]byte(jsonContent), &objs)

	return objs
}

func writeRecordsToFile(records []Record, file string) {
	// open file to write data to
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
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

	// write data to file
	f.Write(data)
	f.Close()
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
