package programs

import (
	"fmt"
)

func Help(in chan string, out chan interface{}, err chan error) {
	args := ExtractArgs(in)

	if len(args) != 0 {
		err <- fmt.Errorf("too many arguments")
		return
	}

	out <- "LIST OF AVAILABLE COMMANDS:\n" +
		"- cat [FILE] - reads file FILE\n" +
		"- cd DIR - travels to DIR\n" +
		"- du - shows disk usage\n" +
		"- ls - shows list of entries in current directory\n" +
		"- mkdir DIR - creates directory in current directory\n" +
		"- rmdir DIR - remove directory from current directory\n" +
		"- tail [-n NUM --lines NUM] FILE - reads last NUM lines of file FILE\n" +
		"- touch FILE - updates FILE modification time. If FILE doesn`t exists, creates new empty file"
	err <- nil
}
