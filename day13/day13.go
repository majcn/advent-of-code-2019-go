package main

import (
	"majcn.si/advent-of-code-2019/intcode"
	. "majcn.si/advent-of-code-2019/util"
	"fmt"
)

func parseData() string {
	return FetchInputData(13)
}

func solvePart1(data string) (rc int) {
	runner := intcode.NewRunner(data)

	inC := make(chan int)
	outC := make(chan int)
	statusC := make(chan intcode.StatusType)

	go runner.Exec(inC, outC, statusC)

	for {
		select {
		case <-outC:
			<-outC
			tile := <-outC

			if tile == 2 {
				rc++
			}
			break
		case status := <-statusC:
			if status == intcode.StatusFinished {
				return
			}
			break
		}
	}
}

func solvePart2(data string) (rc int) {
	runner := intcode.NewRunner(data)
	runner.Program[0] = 2

	inC := make(chan int)
	outC := make(chan int)
	statusC := make(chan intcode.StatusType)

	go runner.Exec(inC, outC, statusC)

	ballX := 0
	paddleX := 0

	for {
		select {
		case x := <-outC:
			y := <-outC
			tile := <-outC

			if x == -1 && y == 0 {
				rc = tile
			} else if tile == 3 {
				paddleX = x
			} else if tile == 4 {
				ballX = x
			}
		case status := <-statusC:
			if status == intcode.StatusNeedInput {
				if paddleX > ballX {
					inC <- -1
				} else if paddleX == ballX {
					inC <- 0
				} else {
					inC <- 1
				}
			}
			if status == intcode.StatusFinished {
				return
			}
		}
	}
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
