package interpreter

import "github.com/tot0p/Ecla/interpreter/eclaType"

type BusType int

const (
	BUS_MAIN BusType = iota
	BUS_RETURN
	BUS_NONE
)

var noneBus Bus = Bus{
	Type: BUS_NONE,
	Val:  nil,
}

type Bus struct {
	Type BusType
	Val  eclaType.Type
}

func (b *Bus) TransformTo(busType BusType) {
	b.Type = busType
}

func (b *Bus) IsReturn() bool {
	return b.Type == BUS_RETURN
}

func (b *Bus) IsMain() bool {
	return b.Type == BUS_MAIN
}

func (b *Bus) IsNone() bool {
	return b.Type == BUS_NONE
}

func (b *Bus) GetVal() eclaType.Type {
	return b.Val
}

func NewMainBus(Val eclaType.Type) *Bus {
	return &Bus{
		Type: BUS_MAIN,
		Val:  Val,
	}
}

func NewReturnBus(Val eclaType.Type) *Bus {
	return &Bus{
		Type: BUS_RETURN,
		Val:  Val,
	}
}

func NewNoneBus() *Bus {
	return &noneBus
}
