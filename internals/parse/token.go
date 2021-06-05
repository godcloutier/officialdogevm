package parse

import "fmt"

type TokenType int

type Token struct {
	Type   TokenType
	Val    string
	Column Pos
	Line   int
}

func (t Token) String() string {
	if t.Type == Semicolon {
		return fmt.Sprintf("[%s, \\n]", TokenTypeMap[Semicolon])
	}

	if t.Type == Eof {
		return fmt.Sprintf("[EOF, End of file]")
	}
	return fmt.Sprintf("[%s, %s]", TokenTypeMap[t.Type], t.Val)
}

// End of file
const eof = -1

const (
	// ILLEGAL Token
	Illegal TokenType = iota

	// Identifiers + literals
	Ident    // add, foobar, x, y, ...
	Int      // 1343456
	String   // "hello world"
	Function // func
	Contract // contract

	IntType
	StringType
	BoolType
	VoidType

	Assign   // =
	Plus     // +
	Minus    // -
	Bang     // !
	Asterisk // *
	Slash    // /
	Mod      // %

	PlusAssign     // +=
	MinusAssign    // -=
	AsteriskAssign // *=
	SlashAssign    // /=
	ModAssign      // %=

	Land // &&
	Lor  // ||
	Inc  // ++
	Dec  //--

	LT     // <
	GT     // >
	LTE    // <=
	GTE    // >=
	EQ     // ==
	NOT_EQ // !=

	Comma // ,

	Lparen // (
	Rparen // )
	Lbrace // {
	Rbrace // }

	True   // true
	False  // false
	If     // if
	Else   // else
	Return // return
	Eof    // end of file
	Eol    // end of line
	Semicolon
)

// TokenTypeMap mapping TokenType with its
// string, this helps debugging
var TokenTypeMap = map[TokenType]string{
	Illegal: "ILLEGAL",

	Ident:    "IDENT",
	Int:      "INT",
	String:   "STRING",
	Function: "FUNCTION",
	Contract: "CONTRACT",

	IntType:    "INT_TYPE",
	StringType: "STRING_TYPE",
	BoolType:   "BOOL_TYPE",

	Assign:   "ASSIGN",
	Plus:     "PLUS",
	Minus:    "MINUS",
	Bang:     "BANG",
	Asterisk: "ASTERISK",
	Slash:    "SLASH",
	Mod:      "MOD",

	PlusAssign:     "PLUS_ASSIGN",
	MinusAssign:    "MINUS_ASSIGN",
	AsteriskAssign: "ASTERISK_ASSIGN",
	SlashAssign:    "SLASH_ASSIGN",
	ModAssign:      "MOD_ASSIGN",

	Land: "LAND",
	Lor:  "LOR",
	Inc:  "INC",
	Dec:  "DEC",

	LT:     "LT",
	GT:     "GT",
	LTE:    "LTE",
	GTE:    "GTE",
	EQ:     "EQ",
	NOT_EQ: "NOT_EQ",

	Comma: "COMMA",

	Lparen: "LPAREN",
	Rparen: "RPAREN",
	Lbrace: "LBRACE",
	Rbrace: "RBRACE",

	True:   "TRUE",
	False:  "FALSE",
	If:     "IF",
	Else:   "ELSE",
	Return: "RETURN",

	Eof:       "EOF",
	Eol:       "EOL",
	Semicolon: "SEMICOLON",
}

var keywords = map[string]TokenType{
	"contract": Contract,
	"func":     Function,
	"if":       If,
	"else":     Else,
	"int":      IntType,
	"string":   StringType,
	"bool":     BoolType,
	"return":   Return,
	"true":     True,
	"false":    False,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}