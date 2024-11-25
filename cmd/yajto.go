package main

import (
	"fmt"

	"github.com/luquor/yajto/internal"
)

func main() {
	inputFile, outputFile, extension, dryRun := internal.ParseArguments()

	fmt.Printf("Input File: %s\n", inputFile)
	fmt.Printf("Output File: %s\n", outputFile)
	fmt.Printf("Extension: %s\n", extension)
	fmt.Printf("Dry Run: %t\n", dryRun)
}
