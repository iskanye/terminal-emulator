package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Исполнить скрипт
func ExecuteScript(script string) {
	var i int = 1
	file, _ := os.Open(script)
	reader := bufio.NewReader(file)

	for {
		input, err := reader.ReadString('\n')
		trimmedInput := strings.TrimSpace(input)

		PrintInputField()
		fmt.Println(trimmedInput)

		parserErr := Parser(trimmedInput)
		if parserErr != nil {
			fmt.Print(fmt.Sprintf("line %d: ", i), parserErr)
			break
		}

		if err != nil {
			break
		}

		i++
	}

	fmt.Print("\n\"" + script + "\" executed")
}
