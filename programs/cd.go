package programs

func Cd(in, out chan string) {
	in <- "cd"
	for i := range out {
		in <- i
	}
}
