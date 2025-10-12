package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Du() {
	args := ExtractArgs(stdin)

	if len(args) > 0 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	for _, i := range vfs.FileExplorer.GetCurrent().Children {
		output := fmt.Sprintf("%6d    %s", i.GetSize(), i.Name)
		stdout <- output
	}

	stderr <- nil
}
