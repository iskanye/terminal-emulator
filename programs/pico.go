package programs

import (
	"fmt"
	"strings"
	"terminal-emulator/vfs"
)

func Pico(in chan string, out chan any, stderr chan error) {
	args := ExtractArgs(in)

	if len(args) > 2 {
		stderr <- fmt.Errorf("too many arguments")
		return
	} else if len(args) < 2 {
		stderr <- fmt.Errorf("not enough arguments")
		return
	}

	file, err := vfs.FileExplorer.GetFile(args[0])
	if err != nil {
		stderr <- err
		return
	}

	file.Content = strings.Trim(args[1], "\"")
	stderr <- nil
}
