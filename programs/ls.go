package programs

func Ls(in, out chan string) {
	in <- "ls"
	for i := range out {
		in <- i
	}
}
