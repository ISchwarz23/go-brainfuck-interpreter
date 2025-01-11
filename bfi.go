package main

import (
	"bfi/memory"
	"bfi/tokenizer"
	"fmt"
	"os"
	"strings"
)

func main() {
	// retrieve arguments
	args := os.Args

	// check arguments
	if len(args) != 2 {
		fmt.Println("Wrong number of arguments were passed.")
		// TODO: print help
		return
	}

	// read the file content
	content, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error on reading File")
	}
	instructions := string(content)
	instructions = strings.ReplaceAll(instructions, "\n", "")
	instructions = strings.ReplaceAll(instructions, "\r", "")

	// tokenize the input
	tokens := tokenizer.Tokenize(instructions)

	// run the bf program
	executeInstructions(memory.New(), tokens, 0)
	fmt.Println()

	// debug output
	// fmt.Println(memory.ToArray())
}

func executeInstructions(memory memory.Memory, tokens []tokenizer.Token, currentInstruction int) int {

	i := currentInstruction
	for i < len(tokens) {
		switch tokens[i] {
		case tokenizer.MOVE_POINTER_RIGHT:
			memory.MovePointerRight()
		case tokenizer.MOVE_POINTER_LEFT:
			memory.MovePointerLeft()
		case tokenizer.INCREMENT_CURRENT_REGISTER:
			memory.GetCurrentRegister().IncrementValue()
		case tokenizer.DECREMENT_CURRENT_REGISTER:
			memory.GetCurrentRegister().DecrementValue()
		case tokenizer.LOOP_START:
			var loopEndI int
			for memory.GetCurrentRegister().GetValue() > 0 {
				loopEndI = executeInstructions(memory, tokens, i+1)
			}
			i = loopEndI
		case tokenizer.LOOP_END:
			return i
		case tokenizer.PRINT_CURRENT_REGISTER:
			fmt.Printf("%c", memory.GetCurrentRegister().GetValue())
		}
		i++
	}

	return i
}
