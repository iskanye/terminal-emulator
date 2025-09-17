package programs

import (
	"fmt"
	"terminal-emulator/vfs"
)

func Du(in chan string, out chan interface{}, err chan error) {
	args := make([]string, 0)
	for i := range in {
		args = append(args, i)
	}

	for _, i := range args {
		err <- fmt.Errorf("unknown argument: %s", i)
		return
	}

	for _, i := range vfs.FileExplorer.GetCurrent().Children {
		output := fmt.Sprintf("%d\t%s", i.GetSize(), i.Name)
		out <- output
	}

	err <- nil
}
