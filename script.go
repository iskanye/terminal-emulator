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
		if err != nil {
			break
		}

		PrintInputField()
		fmt.Print(input)
		Parser(strings.TrimSpace(input[:len(input)-1]))
	}

	fmt.Println("\"" + script + "\" executed")
}
