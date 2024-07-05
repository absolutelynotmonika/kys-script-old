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
   dev_print(fmt.Sprintf("[ TOKEN ] Added new token on line %v (lexeme %v)", line, lexeme))
}

// function to peek at the next value in the input code
func (l* Lexer) Peek() (string, error) {
   if l.Position+1 <= len(l.Code) { 
      return string(l.Code[l.Position+1]), nil
   } else {
      return "", fmt.Errorf("No next character.")
   }
}

func (l* Lexer) Lex() {
   // pre initialised current values neccesary in the code.
   identf_match := regexp.MustCompile(`[a-zA-Z_][a-zA-Z_0-9]*`)
   num_match := regexp.MustCompile(`\d+`)

   // useful information vars
   currentLine := 0
   isAtEnd := false

   // main loop that lexes the code
   for {
      if l.Position >= len(l.Code) { 
         break 
      } else if l.Position == len(l.Code) {
         isAtEnd = true
      }

      // get the current position in the code 
      // as a string
      currentChar := string(l.Code[l.Position])

      // map containing single characters, for checking ofc
      chars := map[string]TokenType{
         "\n": NEWLINE,
         " ": WHITESPACE,
         "(": L_PAREN,
         ")": R_PAREN,
         ".": DOT,
         ",": COLON,
         "+": PLUS,
         "*": STAR,
         "{": L_BRACE,
         "}": R_BRACE,
         "-": MINUS,
      }

      tokType, isChar := chars[currentChar] // check if the char is part of it
      if isChar {
         l.AddToken(currentChar, tokType, currentLine) // add
         l.Advance(1)
         continue
      }

      // double characters
      nextChar, err := l.Peek()

      if !isAtEnd && err == nil {
         if currentChar == "=" && nextChar == "=" {
            l.AddToken("==", DOUBLE_EQUAL, currentLine)
            l.Advance(2)
            continue
         }
      }

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
      fmt.Println("IncurrentCharalid token found")
      l.Advance(1)
   }

   // marking the end of file
   l.AddToken("eof", EOF, currentLine+1)

   dev_print("reached eof")
}
