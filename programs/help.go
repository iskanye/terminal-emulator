package programs

import (
	"fmt"
)

func Help() {
	args := ExtractArgs(stdin)

	if len(args) != 0 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	stdout <- "LIST OF AVAILABLE COMMANDS:\n" +
		"- cat FILE - reads content from FILE\n" +
		"- cd DIR - travels to DIR\n" +
		"- du - shows disk usage\n" +
		"- ls - shows list of entries in current directory\n" +
		"- mkdir DIR - creates directory in current directory\n" +
		"- rmdir DIR - remove directory from current directory\n" +
		"- tail [-n NUM --lines NUM] FILE - reads last NUM lines of file FILE\n" +
		"- touch FILE - updates FILE modification time. If FILE doesn`t exists, creates new empty file\n" +
		"- pico FILE \"CONTENT\" - rewrites FILE content with CONTENT"
	stderr <- nil
}
