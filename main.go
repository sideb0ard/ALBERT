package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mgutz/ansi"
)

func main() {

	PS2 := ansi.Color("#BERT> ", "white")
	setMaxProc()

	shellReader := bufio.NewReader(os.Stdin)
	for {
		// SETUP && LOOP
		fmt.Printf(PS2)
		input, err := shellReader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				myexit()
			} else {
				fmt.Println("Got an err for ye: ", err)
			}
		}

		input = strings.TrimSpace(input)
		response := chompyChomp(input)
		fmt.Println("DUDDDDDE!", response)

	}
}
