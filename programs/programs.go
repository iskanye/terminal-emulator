package programs

type Program struct {
	cmd  string
	exec ProgramFunction
}

type ProgramFunction func(in, out chan string)

var programs []Program = []Program{
	Program{"ls", Ls},
	Program{"cd", Cd},
}
