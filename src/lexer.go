package kysscript

import (
	"fmt"
	"unicode"
)

// useful functions
func Error(line int, message string) {
	fmt.Printf("E: At line %v: %v\n", line, message)
}

func IsDigit(char byte) bool {
	return char <= '0' && char >= '9'
}

// main lexer struct
type Lexer struct {
	SourceCode string
	Tokens []Token
	Position,
	ErrorCount int
}

func (l* Lexer) Advance(times int) {
	l.Position += times
	lexer_dev_print(fmt.Sprintf("[ ADVANCE ] advanced by %v (value: prev. %v, curr. %v).", times, l.Position - times, l.Position))
}

func (l* Lexer) AddToken(lexeme string, toktype TokenType, line int) {
	l.Tokens = append(l.Tokens, Token{
		Lexeme: lexeme,
		Type: toktype,
		Line: line,
	})
	lexer_dev_print(fmt.Sprintf("[ TOKEN ] Added new token on line %v (lexeme \"%v\")", line, lexeme))
}

// function to peek at the next value in the input code
func (l* Lexer) Peek() string {
	if l.Position+1 < len(l.SourceCode) { 
		return string(l.SourceCode[l.Position+1])
	}
	return "\x00" // means theres no next character
}

func (l* Lexer) PeekDouble() byte {
	if l.Position+2 >= len(l.SourceCode) { return '\x00' }
	return l.SourceCode[l.Position+2]
}

func (l* Lexer) GetPattern(start, end int) string {
	return l.SourceCode[start:end]
}

// function to check if the next character is the end of the source code
func (l* Lexer) NextIsEnd() bool {
	return l.Position+1 >= len(l.SourceCode)
}

// function to check if the current character is the end of the source code
func (l* Lexer) IsAtEnd() bool {
	return l.Position >= len(l.SourceCode)
}

func (l* Lexer) IsAlphabetic(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}



func (l* Lexer) Lex() {
	// keywords map
	keywords := map[string]TokenType{
		"if": IF,
		"elseif": ELSEIF,
		"else": ELSE,
		"var": VAR,
		"func": FUNC,
		"true": TRUE,
		"false": FALSE,
		"print": PRINT,
		"return": RETURN,
	}

	// useful information vars
	currentLine := 0

	/////// main loop that lexes the source code
	lexer_dev_print("[ START ] Loop will start")
	for {
		// check if the code ended
		if l.IsAtEnd() { break }

		// get the current position in the source code as a string
		currentChar := string(l.SourceCode[l.Position])
		lexer_dev_print(fmt.Sprintf("[ INFO ] Current character in iteration: %v", currentChar))

		// check if its a single character
		// or operator/double characters
		switch currentChar {
		case " ", "\t":
			l.AddToken(" ", WHITESPACE, currentLine)
			l.Advance(1)
			continue

		case "\n":
			l.AddToken("\n", NEWLINE, currentLine)
			currentLine++
			l.Advance(1)
			continue

		/* Comments can have an ending # or a newline,
		 * But for now, it cant handle multiline comments
		 * directly.
		*/
		case "#":
			for (l.Peek() != "#" || l.Peek() != "\n") && !l.NextIsEnd() {
				l.Advance(1)
			}
			l.Advance(2) // advance the ending # or newline
			continue

		case "\"":
			start := l.Position
			end := l.Position+1

			for l.Peek() != "\"" && !l.IsAtEnd() {
				end++
				if l.Peek() == "\n" { currentLine++ } // natively multiline strings!
				// p.s:  theyre easier to implement than remove.
				l.Advance(1)
			}

			if l.IsAtEnd() {
				Error(currentLine, "Unterminated string.")
				return;
			}

			l.AddToken(l.GetPattern(start+1, end), STRING, currentLine) 
			l.Advance(2) // eat the ending " and yes
			continue

		case "+":
			l.AddToken("+", PLUS, currentLine)
			l.Advance(1)
			continue

		case "-":
			l.AddToken("-", MINUS, currentLine)
			l.Advance(1)
			continue

		case "*":
			l.AddToken("*", STAR, currentLine)
			l.Advance(1)
			continue

		case "/":
			l.AddToken("/", SLASH, currentLine)
			l.Advance(1)
			continue

		case ".":
			l.AddToken(".", DOT, currentLine)
			l.Advance(1)
			continue


		// double chars here
		case ">":
			if l.Peek() == "=" {
				l.AddToken(">=", GREATER_OR_EQUAL, currentLine)
				l.Advance(2)
				continue
			} else {
				l.AddToken(">", GREATER, currentLine)
				l.Advance(1)
				continue
			}

		case "<":
			if l.Peek() == "=" {
				l.AddToken("<=", LESS_OR_EQUAL, currentLine)
				l.Advance(2)
				continue
			} else {
				l.AddToken(">", LESS, currentLine)
				l.Advance(1)
				continue
			}

		case "!":
			if l.Peek() == "=" {
				l.AddToken("!=", NOT_EQUAL, currentLine)
				l.Advance(2)
				continue
			} else {
				l.AddToken("!", NOT, currentLine)
				l.Advance(1)
				continue
			}

		case "=":
			if l.Peek() == "=" {
				l.AddToken("==", DOUBLE_EQUAL, currentLine)
				l.Advance(2)
				continue
			} else {
				l.AddToken("=", EQUAL, currentLine)
				l.Advance(1)
				continue
			}

		default:
			// check is numbet
			if unicode.IsDigit(rune(currentChar[0])) {
				start := l.Position
				end := l.Position+1

				for unicode.IsDigit(rune(l.Peek()[0])) { l.Advance(1); end++ }

				if l.Peek() == "." && unicode.IsDigit(rune(l.PeekDouble())) {
					l.Advance(1) // eat the .
					end++ // ↑
					for unicode.IsDigit(rune(l.Peek()[0])) { end++; l.Advance(1) }
				}

				l.Advance(1) // consume the last number
				l.AddToken(l.GetPattern(start, end), NUMBER, currentLine)
				continue
			}

			if l.IsAlphabetic(rune(currentChar[0])) {
				// if i place those where they technically should, the goto label doesnt work
				// i hate myself.
				start := 0
				end := 0

				// check if it is a keyword
				for key, value := range keywords {
					char_after_keyw := l.Position+len(key)

					// if it matches a keyword and has a space aftwarards, proceed
					if char_after_keyw+1 <= len(l.SourceCode) && l.GetPattern(l.Position, char_after_keyw+1) == key+" " {
						l.AddToken(key, value, currentLine)
						l.Advance(len(key))
						goto end_of_loop // no idea why continue doesnt work but hey.

						// if it matches a keyword and theres no afterward/is eof, proceed
					} else if char_after_keyw == len(l.SourceCode) && l.GetPattern(l.Position, char_after_keyw) == key {
						l.AddToken(key, value, currentLine)
						l.Advance(len(key))

						goto end_of_loop // no idea why continue doesnt work but hey.
					}
				}

				// if it isnt a keyword, its an identifier
				start = l.Position
				end = l.Position+1

				for l.Peek() != " " && l.Peek() != "\x00" && l.IsAlphabetic(rune(l.Peek()[0])) { 
					end++
					l.Advance(1) 
				}

				l.AddToken(l.GetPattern(start, end), IDENTIFIER, currentLine)
				l.Advance(1)
				continue

				end_of_loop:
				continue
			}

			// else...
			l.ErrorCount++
			fmt.Println("Invalid token found")
			l.Advance(1)
		}
	}

	// mark the end of file
	l.AddToken("eof", EOF, currentLine+1)
	lexer_dev_print("reached eof")
}
