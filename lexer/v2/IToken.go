package v2

type ITokenType interface {
	Resolve(l *TLexer)
	Get() []string
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
		&BXOR,
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
		&TCOMMENT,
		//TCOMMENTGROUP,
		//TCOMMENTGROUPEND,
	}
)

var (
	TEXT              = "TEXT"
	STRING            = "STRING"
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
	OR                = "OR"
	AND               = "AND"
	EQUAL             = "EQUAL"
	LPAREN            = "LPAREN"
	RPAREN            = "RPAREN"
	EOL               = "EOL"
	DQUOTE            = "DQUOTE"
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
