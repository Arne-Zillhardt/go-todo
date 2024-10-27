package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetUserInput(question string) (input string) {
	fmt.Print(question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return
}
