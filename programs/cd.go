package programs

func Cd(in, out chan interface{}) {
	out <- "cd"
	for i := range in {
		out <- i
	}
}
