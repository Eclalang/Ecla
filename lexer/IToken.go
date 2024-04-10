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

func findNameInEveryTokenType(name string) ITokenType {
	for _, tokenType := range Every {
		if NameFromGet(tokenType.Get()) == name {
			return tokenType
		}
	}
	return &TokenTypeBaseBehavior{Name: "NULL"}
}

func findNameInTriggerTokenType(name string) *TokenTypeTriggerBehavior {
	for _, tokenType := range Trigger {
		if NameFromGet(tokenType.Get()) == name {
			return tokenType
		}
	}
	return &TokenTypeTriggerBehavior{Name: "NULL"}
}

func findNameInMergerTokenType(name string) *TokenTypeMergerBehavior {
	for _, tokenType := range Merger {
		if NameFromGet(tokenType.Get()) == name {
			return tokenType
		}
	}
	return &TokenTypeMergerBehavior{Name: "NULL"}
}

func findNameInBaseTokenType(name string) *TokenTypeBaseBehavior {
	for _, tokenType := range Base {
		if NameFromGet(tokenType.Get()) == name {
			return tokenType
		}
	}
	return &TokenTypeBaseBehavior{Name: "NULL"}
}

func NameFromGet(Get []string) string {
	return Get[len(Get)-1]
}

func BuildEvery() []ITokenType {
	for _, behavior := range Base {
		Every = append(Every, behavior)
	}
	for _, behavior := range Merger {
		Every = append(Every, behavior)
	}
	for _, behavior := range Trigger {
		Every = append(Every, behavior)
	}
	for _, behavior := range Spaces {
		Every = append(Every, behavior)
	}
	for _, behavior := range Nullable {
		Every = append(Every, behavior)
	}
	return Every
}

var (
	Every []ITokenType             = []ITokenType{}
	Base  []*TokenTypeBaseBehavior = []*TokenTypeBaseBehavior{
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
		&BINT,
		&BADD,
		&BASSIGN,
		&BSUB,
		&BMOD,
		&BDIV,
		&BNOT,
		&BMULT,
		&BLSS,
		&BGTR,
		&BCOMMENTGROUPEND,
		&BTEXT,
	}
	Trigger []*TokenTypeTriggerBehavior = []*TokenTypeTriggerBehavior{
		&TDQUOTE,
		&TSQUOTE,
	}
	Merger []*TokenTypeMergerBehavior = []*TokenTypeMergerBehavior{
		&TCOMMENT,
		&TCOMMENTGROUP,
	}
	Spaces []*TokenTypeSpacesBehavior = []*TokenTypeSpacesBehavior{
		&EMPTY,
		&RETURN,
	}
	Nullable []*TokenTypeNullableBehavior = []*TokenTypeNullableBehavior{
		&NBSLASH,
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
	BSLASH            = "BSLASH"
)
