package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Исполнить скрипт
func ExecuteScript(script string) {
	defer PrintInputField()
	var i int = 1
	file, err := os.Open(script)

	if err != nil {
		Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		PrintInputField()
		Println(input)

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
