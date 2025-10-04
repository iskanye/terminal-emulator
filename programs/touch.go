package programs

import (
	"fmt"
	"time"

	"terminal-emulator/vfs"
)

func Touch(in chan string, out chan any, stderr chan error) {
	args := ExtractArgs(in)

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	} else if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	node, err := vfs.FileExplorer.GetFile(args[0])
	if err != nil {
		err = vfs.FileExplorer.AddNode(args[0], false)
		if err != nil {
			stderr <- fmt.Errorf("directory of same name exists: %s", args[0])
			return
		}
	} else {
		node.Modified = time.Now()
	}

	stderr <- nil
}
