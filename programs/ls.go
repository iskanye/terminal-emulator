package programs

func Ls(in, out chan interface{}) {
	out <- "ls"
	for i := range in {
		out <- i
	}
}
