package lexer

type ITokenType interface {
	// Resolve use the actual l *TLexer to create a new Token. the lexer is ready for the next step when the resolve end
	Resolve(l *TLexer)
	// Get return all the know syntax of an ITokenType, with his name as the last element of the array.
	Get() []string
	// InvolvedWith return all the ITokenType involved with this ITokenType.
	InvolvedWith() []ITokenType
}

// Token is a struct that contains all the information about a token
type Token struct {
	TokenType string
	Value     string
	Position  int
	Line      int
}

func findNameInEveryTokenType(name string, Every []ITokenType) ITokenType {
	for _, tokenType := range Every {
		if tokenType.Get()[len(tokenType.Get())-1] == name {
			return tokenType
		}
	}
	return nil
}

var (
	Every []ITokenType = []ITokenType{
		&BPERIOD,
		&BCOLON,
		&BLBRACE,
		&BRBRACE,
		&BRBRACKET,
		&BLBRACKET,
		&BCOMMA,
		&BBOOL,
		&BXORBIN,
		&BOR,
		&BAND,
		&BLPAREN,
		&BRPAREN,
		&BEOL,
		&BMURLOC,
		&EMPTY,
		&RETURN,
		&BINT,
		&BADD,
		&BASSIGN,
		&BSUB,
		&BMOD,
		&BDIV,
		&BNOT,
		&BMULT,
		&TDQUOTE,
		&TSQUOTE,
		&TCOMMENT,
		&TCOMMENTGROUP,
	}
)

var (
	TEXT              = "TEXT"
	STRING            = "STRING"
	CHAR              = "CHAR"
	PRINT             = "PRINT"
	INT               = "INT"
	FLOAT             = "FLOAT"
	ADD               = "ADD"
	SUB               = "SUB"
	MULT              = "MULT"
	DIV               = "DIV"
	QOT               = "QOT"
	MOD               = "MOD"
	INC               = "INC"
	DEC               = "DEC"
	ASSIGN            = "ASSIGN"
	LSS               = "LSS"
	GTR               = "GTR"
	NEQ               = "NEQ"
	LEQ               = "LEQ"
	GEQ               = "GEQ"
	XOR               = "XOR"
	XORBIN            = "XORBIN"
	OR                = "OR"
	AND               = "AND"
	EQUAL             = "EQUAL"
	LPAREN            = "LPAREN"
	RPAREN            = "RPAREN"
	EOL               = "EOL"
	DQUOTE            = "DQUOTE"
	SQUOTE            = "SQUOTE"
	PERIOD            = "PERIOD"
	COLON             = "COLON"
	LBRACE            = "LBRACE"
	RBRACE            = "RBRACE"
	LBRACKET          = "LBRACKET"
	RBRACKET          = "RBRACKET"
	COMMA             = "COMMA"
	NOT               = "NOT"
	BOOL              = "BOOL"
	MURLOC            = "MURLOC"
	EOF               = "EOF"
	COMMENT           = "COMMENT"
	COMMENTGROUP      = "COMMENTGROUP"
	COMMENTGROUPIDENT = "COMMENTGROUPIDENT"
)
