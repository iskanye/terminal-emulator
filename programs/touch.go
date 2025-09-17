package programs

import (
	"fmt"
	"terminal-emulator/vfs"
	"time"
)

func Touch(in chan string, out chan interface{}, stderr chan error) {
	args := make([]string, 0)
	for i := range in {
		args = append(args, i)
	}

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	} else if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	node, err := vfs.FileExplorer.GetFile(args[0])
	if err != nil {
		vfs.FileExplorer.AddNode(args[0], false)
	} else {
		node.Modified = time.Now()
	}

	stderr <- nil
}
