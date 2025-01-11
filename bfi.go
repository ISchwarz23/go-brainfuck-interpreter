package main

import (
	"bfi/memory"
	"fmt"
	"os"
)

func main() {
	// Retrieve all arguments
	args := os.Args

	// Print additional arguments
	if len(args) != 2 {
		fmt.Println("Wrong number of arguments were passed.")
		// TODO: print help
		return
	}

	bfFilePath := args[1]
	fmt.Println(bfFilePath)

	var memory = memory.New()
	memory.MovePointerRight()
	memory.IncrementValueAtCurrentPointerPosition()
	memory.MovePointerRight()
	memory.MovePointerRight()
	fmt.Println(memory.ToArray())
}
