package main

import (
	"fmt"

	commandline "github.com/michaelchristwin/N_M3U8DL-RE-go.git/app/command_line"
	log "github.com/michaelchristwin/N_M3U8DL-RE-go.git/common/log"
)

func main() {
	options := commandline.CommandInvoker()
	fmt.Printf("Keys: %v\n", *options.Keys)
	console := log.NewCustomAnsiConsole(false, true)
	console.SuccessMessage("Hear me subjects of Ymir")
}
