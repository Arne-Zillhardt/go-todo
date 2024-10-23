package test

import (
	"encoding/csv"
	"log"
	"os"
	"testing"
)

var readFile [][]string

func AssertEquals(expected any, actual any, t *testing.T) {
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
	}
}

func SaveFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	readFile, _ = reader.ReadAll()
}

func WriteSavedFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.WriteAll(readFile)
}
