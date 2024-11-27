package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	file, err := os.ReadFile("./test/samples/sample-01.json")
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// generic map
	var jsonContent map[string]interface{}
	err = json.Unmarshal(file, &jsonContent)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	yamlData, err := yaml.Marshal(jsonContent)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}

	fmt.Println(string(yamlData))
}
