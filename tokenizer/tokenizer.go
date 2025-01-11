package tokenizer

type Token int

const (
	MOVE_POINTER_RIGHT Token = iota
	MOVE_POINTER_LEFT
	INCREMENT_CURRENT_REGISTER
	DECREMENT_CURRENT_REGISTER
	LOOP_START
	LOOP_END
	PRINT_CURRENT_REGISTER
	READ_TO_CURRENT_REGISTER
)

func Tokenize(instructions string) []Token {
	tokens := make([]Token, len(instructions))
	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		case '>':
			tokens[i] = MOVE_POINTER_RIGHT
		case '<':
			tokens[i] = MOVE_POINTER_LEFT
		case '+':
			tokens[i] = INCREMENT_CURRENT_REGISTER
		case '-':
			tokens[i] = DECREMENT_CURRENT_REGISTER
		case '[':
			tokens[i] = LOOP_START
		case ']':
			tokens[i] = LOOP_END
		case '.':
			tokens[i] = PRINT_CURRENT_REGISTER
		case ',':
			tokens[i] = READ_TO_CURRENT_REGISTER
		}
	}
	return tokens
}
