package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Ls() {
	var result []string
	args := ExtractArgs(stdin)

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
		stdout <- i
	}

	stderr <- nil
}
