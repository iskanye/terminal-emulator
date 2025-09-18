package programs

import (
	"fmt"
	"terminal-emulator/vfs"
)

func Ls(in chan string, out chan interface{}, stderr chan error) {
	var result []string
	args := ExtractArgs(in)

	if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	} else if len(args) == 1 {
		var err error
		result, err = vfs.FileExplorer.ListDir(args[0])
		if err != nil {
			stderr <- err
			return
		}
	} else {
		result = vfs.FileExplorer.List()
	}

	for _, i := range result {
		out <- i
	}

	stderr <- nil
}
