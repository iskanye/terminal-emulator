package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Du(in chan string, out chan interface{}, err chan error) {
	args := ExtractArgs(in)

	if len(args) > 0 {
		err <- fmt.Errorf("too many arguments")
		return
	}

	for _, i := range vfs.FileExplorer.GetCurrent().Children {
		output := fmt.Sprintf("%d    %s", i.GetSize(), i.Name)
		out <- output
	}
}
