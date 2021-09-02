package main

import (
	"majcn.si/advent-of-code-2019/intcode"
	. "majcn.si/advent-of-code-2019/util"
	"fmt"
)

func parseData() string {
	return FetchInputData(21)
}

func stringToAsciiIntSlice(s string) []int {
	result := make([]int, len(s))
	for i, c := range s {
		result[i] = int(c)
	}

	return result
}

func solvePartX(data string, commands []string) (rc int) {
	commandsAsIntList := make([]int, 0)
	for _, command := range commands {
		commandsAsIntList = append(commandsAsIntList, stringToAsciiIntSlice(command)...)
		commandsAsIntList = append(commandsAsIntList, 10)
	}

	runner := intcode.NewRunner(data)

	inC := make(chan int)
	outC := make(chan int)
	statusC := make(chan intcode.StatusType)

	go runner.Exec(inC, outC, statusC)

	for {
		select {
		case o := <-outC:
			rc = o
		case status := <-statusC:
			if status == intcode.StatusNeedInput {
				inC <- commandsAsIntList[0]
				commandsAsIntList = commandsAsIntList[1:]
			}
			if status == intcode.StatusFinished {
				return
			}
		}
	}
}

func solvePart1(data string) (rc int) {
	commands := []string{
		"NOT A J",
		"NOT B T",
		"AND D T",
		"OR T J",
		"NOT C T",
		"AND A T",
		"AND B T",
		"AND D T",
		"OR T J",
		"WALK",
	}

	return solvePartX(data, commands)
}

func solvePart2(data string) (rc int) {
	commands := []string{
		"NOT A J",
		"NOT B T",
		"AND D T",
		"OR T J",
		"NOT C T",
		"AND A T",
		"AND B T",
		"AND D T",
		"AND H T",
		"OR T J",
		"RUN",
	}

	return solvePartX(data, commands)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
