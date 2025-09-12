package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const username string = "iskanye"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		defer errorHandler()
		fmt.Print(username + "> ")

		input, _ := reader.ReadString('\n')
		Parser(strings.TrimSpace(input[:len(input)-1]))
	}
}

func errorHandler() {
	if i := recover(); i != nil {
		fmt.Println(i)
		main()
	}
}
