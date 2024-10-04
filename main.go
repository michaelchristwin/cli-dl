package main

import (
	"fmt"

	commandline "github.com/michaelchristwin/N_M3U8DL-RE-go.git/command_line"
)

func main() {
	options := commandline.CommandInvoker()
	fmt.Printf("Keys: %v\n", *options.Keys)
}
