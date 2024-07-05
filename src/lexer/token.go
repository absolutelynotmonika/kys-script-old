package lexer

type TokenType int // fancy way of saying int
const (
   // literal
   IDENTIFIER TokenType = iota
   STRING
   NUMBER
   
   // single char tokens
   L_PAREN
   R_PAREN
   COLON
   DOT
   SEMICOLON
   MINUS
   PLUS
   STAR
   EXCLAMATION
   L_BRACE
   R_BRACE

   // operators and logical and idk
   DOUBLE_EQUAL
   EQUAL

   // keywords
   TRUE
   FALSE
   AND
   OR
   NOT
   FUNC
   IF
   NULL
   ELSE
   ELIF
   RETURN
   VAR
   WHILE
   PRINT

   // idk tbf
   NEWLINE
   WHITESPACE
   EOF
)

type Token struct {
   Lexeme string
   Type TokenType
   Line int
}
