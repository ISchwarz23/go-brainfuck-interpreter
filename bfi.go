package main

import (
	"bfi/memory"
	"bfi/tokenizer"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	// check arguments
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Wrong number of arguments were passed.")
		fmt.Println()
		printHelp()
		return
	}

	// Compile the regex
	re, err := regexp.Compile(`^[\.,\[\]<>+-]+$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	var instructions string
	if re.MatchString(args[1]) {
		// use arg as instructions
		instructions = args[1]
	} else {
		// read instructions from file
		fileContent, err := os.ReadFile(args[1])
		if err != nil {
			fmt.Println("Error on reading File")
			return
		}
		instructions = string(fileContent)
		instructions = strings.ReplaceAll(instructions, "\n", "")
		instructions = strings.ReplaceAll(instructions, "\r", "")
	}

	// tokenize the input
	tokens := tokenizer.Tokenize(instructions)

	// run the bf program
	memory := memory.New()
	executeInstructions(memory, tokens, 0)
	fmt.Println()

	// debug output
	// fmt.Println(memory.ToArray())
}

func executeInstructions(memory memory.Memory, tokens []tokenizer.Token, startInstructionIndex int) int {

	instructionIndex := startInstructionIndex
	for instructionIndex < len(tokens) {
		switch tokens[instructionIndex] {
		case tokenizer.MOVE_POINTER_RIGHT:
			memory.MovePointerRight()
		case tokenizer.MOVE_POINTER_LEFT:
			memory.MovePointerLeft()
		case tokenizer.INCREMENT_CURRENT_REGISTER:
			memory.GetCurrentRegister().IncrementValue()
		case tokenizer.DECREMENT_CURRENT_REGISTER:
			memory.GetCurrentRegister().DecrementValue()
		case tokenizer.LOOP_START:
			var loopEndInstructionIndex int
			for memory.GetCurrentRegister().GetValue() > 0 {
				loopEndInstructionIndex = executeInstructions(memory, tokens, instructionIndex+1)
			}
			instructionIndex = loopEndInstructionIndex
		case tokenizer.LOOP_END:
			return instructionIndex
		case tokenizer.PRINT_CURRENT_REGISTER:
			fmt.Printf("%c", memory.GetCurrentRegister().GetValue())
		case tokenizer.READ_TO_CURRENT_REGISTER:
			memory.GetCurrentRegister().SetValue(readNumber())
		}
		instructionIndex++
	}

	return instructionIndex
}

func readNumber() int {
	var number int
	for {
		fmt.Print("Enter a number: ")
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			// Clear the input buffer to avoid infinite loop due to lingering invalid input
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}
	return number
}

func printHelp() {
	fmt.Println("Interpreter of the brainfuck language.")
	fmt.Println("Example Usages:")
	fmt.Println("  ./bfi hello-world.bf")
	fmt.Println("  ./bfi ,>,<[>+<-]")
}
