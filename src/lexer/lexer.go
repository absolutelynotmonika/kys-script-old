package lexer

import (
   "fmt"
   "regexp"
)

type Error struct {
   Code int
   Message string
}

type Lexer struct {
   Code string
   Tokens []Token
   Position,
   ErrorCount int
}

func (l* Lexer) Advance(times int) {
   l.Position += times
   dev_print(fmt.Sprintf("[ ADVANCE ] advanced by %v (value: prev. %v, curr. %v).", times, l.Position - times, l.Position))
}

func (l* Lexer) AddToken(lexeme string, toktype TokenType, line int) {
   l.Tokens = append(l.Tokens, Token{
      Lexeme: lexeme,
      Type: toktype,
      Line: line,
   })
   dev_print(fmt.Sprintf("[ TOKEN ] Added new token on line %v (lexeme \"%v\")", line, lexeme))
}

// function to peek at the next value in the input code
func (l* Lexer) Peek() string {
   if l.Position+1 < len(l.Code) { 
      return string(l.Code[l.Position+1])
   } else {
      return "" // means theres no next character
   }
}

func (l* Lexer) NextIsEnd() bool {
   return l.Position+1 >= len(l.Code)
}

func (l* Lexer) IsAtEnd() bool {
   return l.Position >= len(l.Code)
}

func (l* Lexer) Lex() {
   ////// pre initialised current values neccesary in the code.
   identf_match := regexp.MustCompile(`[a-zA-Z_][a-zA-Z_0-9]*`)
   num_match := regexp.MustCompile(`\d+`)
   
   // useful information vars
   currentLine := 0

   /////// main loop that lexes the coded
   dev_print("[ START ] loop will start")
   for {
      // check if the code ended
      if l.IsAtEnd() { break }

      // get the current position in the code 
      // as a string
      currentChar := string(l.Code[l.Position])
      dev_print(fmt.Sprintf("[ INFO ] Current character in iteration: %v", currentChar))
  
      // check if its a single character
      // or operator/double characte

      switch currentChar {
      case " ", "\t":
         l.AddToken(" ", WHITESPACE, currentLine)
         l.Advance(1)
         continue

      case "\n":
         l.AddToken("\n", NEWLINE, currentLine)
         l.Advance(1)
         continue

      case "#":
         for l.Peek() != "#" && !l.NextIsEnd() {
            l.Advance(1)
         }
         l.Advance(2) // so the character wont overlap and shit
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
      } // end of statement

      // pattern match for num
      if num_match.MatchString(currentChar) {
         l.AddToken(currentChar, NUMBER, currentLine)
         l.Advance(1)
         continue
      }


      // pattern match for identifier
      if identf_match.MatchString(currentChar) {
         l.AddToken(currentChar, IDENTIFIER, currentLine)
         l.Advance(1)
         continue
      }

      // else...
      l.ErrorCount++
      fmt.Println("Invalid token found")
      l.Advance(1)
   }

   // mark the end of file
   l.AddToken("eof", EOF, currentLine+1)

   dev_print("reached eof")
}
