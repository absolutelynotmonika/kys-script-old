package main

import (
   "kysscript/src/lexer"
	"bufio"
	"fmt"
   "os"
)

// for now this is just a REPL

func main() {
   l := lexer.Lexer {
      Tokens: make([]lexer.Token, 0),
      Position: 0,
   }

   reader := bufio.NewReader(os.Stdin)
   for {
      fmt.Print(">>> ")
      input, err := reader.ReadString('\n')

      if err != nil {
         fmt.Println("\nError when reading REPL input:", err.Error())
         os.Exit(1)
      }

      input = string(input[0:len(input)-1])

      l.Code = input
      l.Lex()
   
      for i, v := range l.Tokens {
         fmt.Printf("token %v at line %v: lexeme - \"%v\", type - %v\n", i+1, v.Line, v.Lexeme, v.Type)
      }
      fmt.Printf("Errors: %v\n", l.ErrorCount)

      l = lexer.Lexer {
         Tokens: make([]lexer.Token, 0),
         Position: 0,
      }
   }
}
