package file

import (
	"encoding/csv"
	"log"
	"os"
)

var readFile [][]string
var filePath string = "./assets/todo.csv"

func GetReadFile() [][]string {
	if readFile == nil {
		readFile = readTheTodoFile()
	}

	return readFile
}

func WriteNewTaskToTheTodoFile(contentToWrite []string) {
	file := openFileToAppend()
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	updateReadFile(contentToWrite)
	writeToFile(*writer, contentToWrite)
}

func UpdateTaskInTheTodoFile(updatedContent []string) {
	file := openFileToRewrite()
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	updateReadFile(updatedContent)
	for _, line := range readFile {
		writeToFile(*writer, line)
	}
}

func RemoveTaskInTheTodoFile(id int) {
	var newFile [][]string
	for i, line := range readFile {
		if i == id {
			continue
		}

		newFile = append(newFile, line)
	}

	readFile = newFile

	file := openFileToRewrite()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, line := range readFile {
		writeToFile(*writer, line)
	}
}

func updateReadFile(updatedContent []string) {
	for i, line := range readFile {
		if line[0] != updatedContent[0] {
			continue
		}

		readFile[i][1] = updatedContent[1]
		readFile[i][2] = updatedContent[2]
		readFile[i][3] = updatedContent[3]
		return
	}

	readFile = append(readFile, updatedContent)
}

func readTheTodoFile() (readContent [][]string) {
	file := openFileToRead()
	defer file.Close()

	reader := csv.NewReader(file)
	readContent, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func writeToFile(writer csv.Writer, lineToWrite []string) {
	err := writer.Write(lineToWrite)
	if err != nil {
		log.Fatal(err)
	}
}

func openFileToRead() *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func openFileToAppend() *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func openFileToRewrite() *os.File {
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
