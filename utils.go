package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mgutz/ansi"
)

var lime = ansi.ColorCode("green+h:black")
var reset = ansi.ColorCode("reset")

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func myexit() {
	fmt.Println("Later, dude...")
	os.Exit(0)
}

func setMaxProc() {
	// Check GOMAXPROCS
	max_procs := runtime.GOMAXPROCS(-1)
	num_procs_to_set := runtime.NumCPU()
	if max_procs != num_procs_to_set {
		runtime.GOMAXPROCS(num_procs_to_set)
	}
}
