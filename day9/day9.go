package main

import (
	"majcn.si/advent-of-code-2019/intcode"
	. "majcn.si/advent-of-code-2019/util"
	"fmt"
)

func parseData() string {
	return FetchInputData(9)
}

func solvePart1(data string) (rc int) {
	runner := intcode.NewRunner(data)

	inC := make(chan int)
	outC := make(chan int)
	statusC := make(chan intcode.StatusType, 2)

	go runner.Exec(inC, outC, statusC)

	inC <- 1
	for c := range outC {
		rc = c
	}

	return
}

func solvePart2(data string) (rc int) {
	runner := intcode.NewRunner(data)

	inC := make(chan int)
	outC := make(chan int)
	statusC := make(chan intcode.StatusType, 2)

	go runner.Exec(inC, outC, statusC)

	inC <- 2
	for c := range outC {
		rc = c
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
