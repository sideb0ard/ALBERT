package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func main() {

	PS2 := ansi.Color("#BERT> ", "white")
	setMaxProc()

	albert := newAgent()

	shellReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(PS2)
		input, err := shellReader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				myexit()
			} else {
				fmt.Println("Got an err for ye: ", err)
			}
		}

		fmt.Println(albert.respond(input))

	}
}
