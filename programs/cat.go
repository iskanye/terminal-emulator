package programs

import (
	"fmt"
	"strings"

	"terminal-emulator/vfs"
)

func Cat() {
	args := ExtractArgs(stdin)

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
				for _, i := range strings.Split(strings.TrimSpace(content), "\n") {
					stdout <- i
				}
			}
		}
	}

	stderr <- nil
}
