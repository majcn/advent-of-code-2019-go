package intcode

import (
	. "majcn.si/advent-of-code-2019/util"
	"strconv"
	"strings"
)

type StatusType int

const (
	StatusNeedInput StatusType = 2
	StatusFinished  StatusType = 99
)

type Runner struct {
	Program map[int]int

	Position     int
	RelativeBase int
}

func (runner *Runner) getOpcode() int {
	return runner.Program[runner.Position] % 100
}

func (runner *Runner) getAddress(positionOffset int) int {
	mode := (runner.Program[runner.Position] / 10 / PowInt(10, positionOffset)) % 10
	switch mode {
	case 1:
		return runner.Position + positionOffset
	case 2:
		return runner.RelativeBase + runner.Program[runner.Position+positionOffset]
	default:
		return runner.Program[runner.Position+positionOffset]
	}
}

func (runner *Runner) getValue(positionOffset int) int {
	return runner.Program[runner.getAddress(positionOffset)]
}

func (runner *Runner) setValue(positionOffset int, value int) {
	runner.Program[runner.getAddress(positionOffset)] = value
}

func (runner *Runner) executeOpcode1() {
	p1 := runner.getValue(1)
	p2 := runner.getValue(2)
	runner.setValue(3, p1+p2)
	runner.Position += 3
}

func (runner *Runner) executeOpcode2() {
	p1 := runner.getValue(1)
	p2 := runner.getValue(2)
	runner.setValue(3, p1*p2)
	runner.Position += 3
}

func (runner *Runner) executeOpcode3(inputValue int) {
	runner.setValue(1, inputValue)
	runner.Position += 1
}

func (runner *Runner) executeOpcode4(out chan<- int) {
	p1 := runner.getValue(1)
	out <- p1
	runner.Position += 1
}

func (runner *Runner) executeOpcode5() {
	p1 := runner.getValue(1)
	p2 := runner.getValue(2)
	if p1 != 0 {
		runner.Position = p2 - 1
	} else {
		runner.Position += 2
	}
}

func (runner *Runner) executeOpcode6() {
	p1 := runner.getValue(1)
	p2 := runner.getValue(2)
	if p1 == 0 {
		runner.Position = p2 - 1
	} else {
		runner.Position += 2
	}
}

func (runner *Runner) executeOpcode7() {
	p1 := runner.getValue(1)
	p2 := runner.getValue(2)
	if p1 < p2 {
		runner.setValue(3, 1)
	} else {
		runner.setValue(3, 0)
	}
	runner.Position += 3
}

func (runner *Runner) executeOpcode8() {
	p1 := runner.getValue(1)
	p2 := runner.getValue(2)
	if p1 == p2 {
		runner.setValue(3, 1)
	} else {
		runner.setValue(3, 0)
	}
	runner.Position += 3
}

func (runner *Runner) executeOpcode9() {
	p1 := runner.getValue(1)
	runner.RelativeBase += p1
	runner.Position += 1
}

func (runner *Runner) Exec(inC <-chan int, outC chan<- int, statusC chan<- StatusType) {
	defer close(outC)
	defer close(statusC)

	for {
		opcode := runner.getOpcode()
		switch opcode {
		case 1:
			runner.executeOpcode1()
			break
		case 2:
			runner.executeOpcode2()
			break
		case 3:
			statusC <- StatusNeedInput
			inputValue := <-inC
			runner.executeOpcode3(inputValue)
			break
		case 4:
			runner.executeOpcode4(outC)
			break
		case 5:
			runner.executeOpcode5()
			break
		case 6:
			runner.executeOpcode6()
			break
		case 7:
			runner.executeOpcode7()
			break
		case 8:
			runner.executeOpcode8()
			break
		case 9:
			runner.executeOpcode9()
			break
		case 99:
			statusC <- StatusFinished
			return
		}

		runner.Position++
	}
}

func NewRunner(data string) Runner {
	dataSplit := strings.Split(data, ",")

	program := make(map[int]int, len(dataSplit))
	for i, v := range dataSplit {
		program[i], _ = strconv.Atoi(v)
	}

	return Runner{Program: program}
}
