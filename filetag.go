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

	for i, record := range records {
		if record.FileName == os.Args[1] {
			existing = i
		}
	}

	if existing < 0 {
		newFile := Record{os.Args[1], os.Args[2:]}
		records = append(records, newFile)
	} else {
		records[existing].Tags =
			append(records[existing].Tags, os.Args[2:]...)
	}

	writeRecordsToFile(records, "data.json")
	fmt.Println("Done...")
}

func getRecordsFromFile(file string) []Record {
	// read in file
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
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
		fmt.Println("Can't data open file...")
		return
	}
	defer f.Close()

	// marshal files to json
	data, err := json.Marshal(records)
	if err != nil {
		fmt.Println("Can't create json...")
	}

	// write data to file
	f.Write(data)
	f.Close()
}
