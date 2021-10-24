package util

import (
	"log"
	"os"
)

// GetOutputDir returns the cirrent directory as the default target
// directory for the diagram will output to
func GetDefaultOutputDir() string {

	outputDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to determine the current directory")
	}

	return outputDir
}
