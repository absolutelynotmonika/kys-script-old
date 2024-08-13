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
	SLASH
	EXCLAMATION
	L_BRACE
	R_BRACE

	// operators and logical and idk
	DOUBLE_EQUAL
	LESS_OR_EQUAL
	GREATER_OR_EQUAL
	NOT_EQUAL
	EQUAL
	LESS
	GREATER
	NOT

	// keywords
	TRUE
	FALSE
	FUNC
	IF
	NULL
	ELSE
	ELSEIF
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
