package interpreter

import "github.com/Eclalang/Ecla/interpreter/eclaType"

// BusType is the type of bus.
type BusType int

const (
	BUS_MAIN   BusType = iota // Main bus
	BUS_RETURN                // Return bus
	BUS_NONE                  // None bus
)

// noneBus is the Global none bus.
var noneBus Bus = Bus{
	Type: BUS_NONE,
	Val:  nil,
}

type Bus struct {
	Type BusType
	Val  eclaType.Type
}

// TransformTo transforms the bus to the given type.
func (b *Bus) TransformTo(busType BusType) {
	b.Type = busType
}

// IsReturn returns true if the bus is a return bus.
func (b *Bus) IsReturn() bool {
	return b.Type == BUS_RETURN
}

// IsMain returns true if the bus is a main bus.
func (b *Bus) IsMain() bool {
	return b.Type == BUS_MAIN
}

// IsNone returns true if the bus is a none bus.
func (b *Bus) IsNone() bool {
	return b.Type == BUS_NONE
}

// GetVal returns the value of the bus.
func (b *Bus) GetVal() eclaType.Type {
	return b.Val
}

// NewMainBus returns a new main bus.
func NewMainBus(Val eclaType.Type) *Bus {
	return &Bus{
		Type: BUS_MAIN,
		Val:  Val,
	}
}

// NewReturnBus returns a new return bus.
func NewReturnBus(Val eclaType.Type) *Bus {
	return &Bus{
		Type: BUS_RETURN,
		Val:  Val,
	}
}

// NewNoneBus returns the Global none bus.
func NewNoneBus() *Bus {
	return &noneBus
}

// IsMultipleBus returns true if the bus is a multiple bus.
func IsMultipleBus(b []*Bus) bool {
	return len(b) > 1
}
