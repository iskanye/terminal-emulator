package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ExecuteScript(script string) {
	file, _ := os.Open(script)
	reader := bufio.NewReader(file)

	for {
		input, err := reader.ReadString('\n')
		trimmedInput := strings.TrimSpace(input)

		PrintInputField()
		fmt.Println(trimmedInput)
		Parser(trimmedInput)

		if err != nil {
			break
		}
	}

	fmt.Println("\"" + script + "\" executed")
}
