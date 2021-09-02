package main

import (
	"majcn.si/advent-of-code-2019/intcode"
	. "majcn.si/advent-of-code-2019/util"
	"bufio"
	"fmt"
	"os"
)

func parseData() string {
	return FetchInputData(25)
}

func stringToAsciiIntSlice(s string) []int {
	result := make([]int, len(s))
	for i, c := range s {
		result[i] = int(c)
	}

	return result
}

func solvePart1(data string) (rc int) {
	runner := intcode.NewRunner(data)
	reader := bufio.NewReader(os.Stdin)

	inC := make(chan int, 20)
	outC := make(chan int)
	statusC := make(chan intcode.StatusType)

	go runner.Exec(inC, outC, statusC)

	for {
		select {
		case o := <-outC:
			print(string(byte(o)))
		case status := <-statusC:
			if status == intcode.StatusNeedInput {
				if len(inC) > 0 {
					break
				}

				char, _, _ := reader.ReadRune()
				var command string
				switch char {
				case 'w':
					command = "north"
					break
				case 'a':
					command = "west"
					break
				case 's':
					command = "south"
					break
				case 'd':
					command = "east"
					break
				}

				for _, i := range stringToAsciiIntSlice(command) {
					inC <- i
				}
				inC <- 10
			}
		}
	}
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
}
