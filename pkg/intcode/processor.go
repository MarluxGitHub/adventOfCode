package intcode

type Interpreter struct {
	Instructions []int
}

func (i *Interpreter) Run() {
	for j := 0; j < len(i.Instructions); j += 4 {
		opcode := i.Instructions[j]
		if opcode == 99 {
			break
		}

		input1 := i.Instructions[i.Instructions[j+1]]
		input2 := i.Instructions[i.Instructions[j+2]]
		output := i.Instructions[j+3]

		if opcode == 1 {
			i.Instructions[output] = input1 + input2
		} else if opcode == 2 {
			i.Instructions[output] = input1 * input2
		}
	}
}

func (i *Interpreter) GetValueOfRegister(register int) int {
	return i.Instructions[register]
}

func NewInterpreter(instructions []int) *Interpreter {
	return &Interpreter{Instructions: instructions}
}
