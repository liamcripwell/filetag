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

/*
TODO:
 - get records with tag
 - list tags for specific file
*/
func main() {
	// get command argument
	if len(os.Args) > 1 {
		command := os.Args[1]

		// read existing records from disk
		records := getRecordsFromFile("data.json")

		// perform chosen command
		switch command {
		case "tag":
			addContent(records)
		case "untag":
			removeContent(records)
		case "list":
			// check if a filename is given
			if len(os.Args) > 2 {
				listTags(records, os.Args[2])
			} else {
				fmt.Println("no filename provided...")
			}
		case "with":
			listRecords(records)
		}
	} else {
		fmt.Println("no command provided...")
	}
}

// Lists all files which possess all specified tags
func listRecords(records []Record) {
	for _, record := range records {
		// identify matching files
		match := 0
		for _, tag := range os.Args[2:] {
			if stringInSlice(tag, record.Tags) {
				match++
			} else {
				break
			}
		}
		// output matching files
		if match == len(os.Args[2:]) {
			fmt.Println(record.FileName)
		}
	}
}

// Lists all tags of a specified file
func listTags(records []Record, filename string) {
	for _, record := range records {
		if record.FileName == filename {
			for _, tag := range record.Tags {
				fmt.Print(tag, " ")
			}
			fmt.Println("")
			break
		}
	}
}

// Removes existing content from the data store
func removeContent(records []Record) {
	// TODO: warning

	// if no specific tags have been listed for removal
	if len(os.Args) == 3 {
		for i, record := range records {
			// remove entire record of specified file
			if record.FileName == os.Args[2] {
				records = append(records[:i], records[i+1:]...)
				break
			}
		}
	} else {
		// if specific tags have been listed for removal
		for i := 0; i < len(records); i++ {
			// locate specified file record
			if records[i].FileName == os.Args[2] {
				for j := 0; j < len(records[i].Tags); j++ {
					for _, toRemove := range os.Args[3:] {
						// if a matching tag is found, remove it
						if records[i].Tags[j] == toRemove {
							if len(records[i].Tags) > 1 {
								records[i].Tags =
									append(records[i].Tags[:j], records[i].Tags[j+1:]...)
							} else {
								records[i].Tags = nil
							}
						}
					}
				}
			}
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
		// generate list of unique tags to add
		var newTags []string
		for _, tag := range os.Args[3:] {
			if !stringInSlice(tag, newTags) {
				newTags = append(newTags, tag)
			}
		}
		// create and save new Record
		newFile := Record{os.Args[2], newTags}
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
