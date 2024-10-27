package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetUserInput(question string) (input string) {
	fmt.Print(question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Split(input, "\n")[0]

	return
}
