package main

import (
	"bufio"
	"fmt"
	"goyacc-sample/calc"
	"os"
)

func main() {

	// calc.CalcParse(&calc.CalcLex{S: "aa=1\n"})
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			calc.CalcParse(&calc.CalcLex{S: eqn})
		} else {
			break
		}
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
