package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var username string = "iskanye"
var vfs string = ""
var startScript string = "start"

func main() {
	fmt.Println("Welcome to terminal emulator! (~by iskanye~)")
	ExecuteScript(startScript)
	terminal()
}

func terminal() {
	reader := bufio.NewReader(os.Stdin)

	for {
		PrintInputField()

		input, _ := reader.ReadString('\n')
		err := Parser(strings.TrimSpace(input[:len(input)-1]))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func PrintInputField() {
	fmt.Print(username + "> ")
}
