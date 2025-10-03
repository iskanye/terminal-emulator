package programs

import (
	"fmt"
	"strings"

	"terminal-emulator/vfs"
)

func Cat(in chan string, out chan interface{}, stderr chan error) {
	args := ExtractArgs(in)

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	} else {
		for _, i := range args {
			file, err := vfs.FileExplorer.GetFile(i)
			if err != nil {
				stderr <- err
				return
			}

			content, err := file.Read()
			if err != nil {
				stderr <- err
				return
			}

			if content != "" {
				out <- strings.TrimSpace(content)
			}
		}
	}
}
