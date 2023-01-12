package libs

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
)

type Console struct {
}

func (c *Console) Call(name string, args ...any) eclaType.Type {
	switch name {
	case "print":
		for _, arg := range args {
			switch arg.(type) {
			case eclaType.Type:
				fmt.Print(arg.(eclaType.Type).String())
			default:
				fmt.Print(arg)
			}
		}
	}
	return eclaType.Empty{}
}
