package day7

type Program struct {
	Name string
	Weight int
	ProgramsAbove []string
}

func SolvePart1(programs []Program) Program {
	for _, program := range programs {
		if !isProgramAboveOtherProgram(program.Name, programs) {
			return program
		}
	}

	panic("bottom program not found")
}

func isProgramAboveOtherProgram(programName string, programs []Program) bool {
	for _, program:= range programs {
		for _, programAbove := range program.ProgramsAbove {
			if programAbove == programName {
				return true
			}
		}
	}

	return false
}
