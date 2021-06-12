package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string	
}

const (
	ILLEGAL	= "ILLEGAL"
	EOF		= "EOF"
	
	
	//identifier
	IDENT 	= "IDENT" //add, foobar, x, y, 
	INT		= "INT"  //123456
	
	//enzansi
	ASSIGN	= "="
	PLUS	  = "+"
	
	//delimiter
	COMMA = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	//keyword
	FUNCTION = "FUCNTION"
	LET	= "LET"
	
)
