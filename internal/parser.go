package internal

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ParseArguments() (inputFilepath string, outputFile string, extension string, dryRun bool) {
	outputFilepath := flag.String("o", "", "Path to the output file.")
	extensionFlag := flag.String("e", "json", "Wanted conversion extension file (e.g., 'json').")
	dryRunFlag := flag.Bool("d", false, "Use the dry-run mode; this will print the output file to the standard output.")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Error: Input file is required.")
		fmt.Println("Usage: yajto <input file> [-o <output file>] [-e <extension>] [-d]")
		flag.Usage()
		os.Exit(1)
	}

	inputFilepath = args[0]

	extension = strings.TrimPrefix(*extensionFlag, ".")

	if *outputFilepath != "" {
		outputFile = *outputFilepath
	} else {
		outputFile = strings.TrimSuffix(inputFilepath, filepath.Ext(inputFilepath)) + "." + extension
	}

	dryRun = *dryRunFlag

	return inputFilepath, outputFile, extension, dryRun
}
