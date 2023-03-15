package interpreter

import "github.com/tot0p/Ecla/interpreter/eclaType"

type BusType int

const (
	BUS_MAIN BusType = iota
	BUS_RETURN
)

type Bus struct {
	Type BusType
	Val  *eclaType.Type
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

func (b *Bus) GetVal() *eclaType.Type {
	return b.Val
}

func NewBusMain(Val *eclaType.Type) *Bus {
	return &Bus{
		Type: BUS_MAIN,
		Val:  Val,
	}
}

func NewReturnBus(Val *eclaType.Type) *Bus {
	return &Bus{
		Type: BUS_RETURN,
		Val:  Val,
	}
}
