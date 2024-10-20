package file

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadFile() (readContent [][]string) {
	file, err := os.Open("../../../../assets/todo.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		readContent = append(readContent, line)
	}

	return
}

func WriteNewTaskToFile(contentToWrite []string) {
	file, err := os.OpenFile("../../../../assets/todo.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{})
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Write(contentToWrite)
	if err != nil {
		log.Fatal(err)
	}
}
