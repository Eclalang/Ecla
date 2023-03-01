package parser

const (
	ASSIGN     = "="
	INCREMENT  = "++"
	DECREMENT  = "--"
	ADDASSIGN  = "+="
	SUBASSIGN  = "-="
	DIVASSIGN  = "/="
	MODASSIGN  = "%="
	QOTASSIGN  = "//="
	MULTASSIGN = "*="
)

var (
	AssignOperators = map[string]interface{}{
		ASSIGN:     nil,
		INCREMENT:  nil,
		DECREMENT:  nil,
		ADDASSIGN:  nil,
		SUBASSIGN:  nil,
		DIVASSIGN:  nil,
		MODASSIGN:  nil,
		QOTASSIGN:  nil,
		MULTASSIGN: nil,
	}
)
