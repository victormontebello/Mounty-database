package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type record map[string]interface{} // Interface to hold various data types

var dataFile = "data.json" // File to store data

// LoadData reads data from the file on startup
func LoadData() (map[string]record, error) {
    data := make(map[string]record)
    _, err := os.Stat(dataFile)
    if err != nil && os.IsNotExist(err) {
        return data, nil // File doesn't exist, create an empty map
    }
    if err != nil {
        return nil, fmt.Errorf("error checking data file: %w", err)
    }
    fileData, err := ioutil.ReadFile(dataFile)
    if err != nil {
        return nil, fmt.Errorf("error reading data file: %w", err)
    }
    err = json.Unmarshal(fileData, &data)
    if err != nil {
        return nil, fmt.Errorf("error unmarshalling data: %w", err)
    }
    return data, nil
}

// SaveData writes data to the file on program exit
func SaveData(data map[string]record) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return fmt.Errorf("error marshalling data: %w", err)
    }
    err = ioutil.WriteFile(dataFile, jsonData, 0644) // Adjust permissions as needed
    if err != nil {
        return fmt.Errorf("error writing data to file: %w", err)
    }
    return nil
}

// Insert adds a new record to the in-memory data and saves it to the file
func Insert(key string, record record) error {
    data, err := LoadData() // Load existing data
    if err != nil {
        return err
    }
    data[key] = record
    err = SaveData(data) // Save updated data
    if err != nil {
        return err
    }
    return nil
}

// Delete removes a record from the in-memory data and saves it to the file
func Delete(key string) error {
    data, err := LoadData() // Load existing data
    if err != nil {
        return err
    }
    if _, ok := data[key]; !ok {
        return fmt.Errorf("key '%s' not found", key)
    }
    delete(data, key)
    err = SaveData(data) // Save updated data
    if err != nil {
        return err
    }
    return nil
}

// Select retrieves a record by key from the in-memory data
func Select(key string) (record, error) {
    data, err := LoadData() // Load existing data
    if err != nil {
        return nil, err
    }
    if record, ok := data[key]; ok {
        return record, nil
    }
    return nil, fmt.Errorf("key '%s' not found", key)
}

func main() {
    // Load data on startup
    data, err := LoadData()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Parse command-line arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./main <function> <key> [<value1>=<value2>...]")
        return
    }

    function := os.Args[1]
    key := os.Args[2]
    args := os.Args[3:]

    switch function {
    case "Insert":
        if len(args) < 1 {
            fmt.Println("Usage: Insert <key> <value1>=<value2>...")
            return
        }
        record := parseRecord(args)
        err = Insert(key, record)
    case "Delete":
        err = Delete(key)

	case "GetAll":
		fmt.Println(data)
    default:
        fmt.Println("Unknown function:", function)
    }

    if err != nil {
        fmt.Println(err)
    }
}

// parseRecord parses a list of key=value pairs into a record map
func parseRecord(args []string) map[string]interface{} {
    record := make(map[string]interface{})
    for _, arg := range args {
        keyVal := strings.SplitN(arg, "=", 2)
        if len(keyVal) != 2 {
            fmt.Println("Invalid key-value pair format:", arg)
            return nil
        }
        record[keyVal[0]] = keyVal[1]
    }
    return record
}


