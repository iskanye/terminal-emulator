package main

import (
	"bufio"
	"fmt"
	"os"
)

// Исполнить скрипт
func ExecuteScript(script string) {
	var i int = 1
	file, err := os.Open(script)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		input, err := reader.ReadString('\n')

		PrintInputField()
		Print(input)

		parserErr := Parser(input)
		if parserErr != nil {
			Println(fmt.Sprintf("line %d: ", i) + " " + fmt.Sprint(parserErr))
			break
		}

		if err != nil {
			break
		}

		i++
	}

	Println("\"" + script + "\" executed")
}
