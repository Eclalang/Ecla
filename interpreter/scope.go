package interpreter

import "github.com/tot0p/Ecla/interpreter/eclaType"

type Scope struct {
	Var      map[string]*eclaType.Var
	next     *Scope
	previous *Scope
}

func NewScopeMain() *Scope {
	return &Scope{
		Var:      make(map[string]*eclaType.Var),
		next:     nil,
		previous: nil,
	}
}
func (s *Scope) Set(name string, value *eclaType.Var) {
	cursor := s
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.Var[name] = value
}

func (s *Scope) Get(name string) (*eclaType.Var, bool) {
	cursor := s
	for cursor.next != nil {
		cursor = cursor.next
	}
	var ok bool = false
	var v *eclaType.Var
	for cursor != nil && !ok {
		v, ok = cursor.Var[name]
		cursor = cursor.previous
	}
	return v, ok
}

func (s *Scope) GoDeep() {
	cursor := s
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = &Scope{
		Var:      make(map[string]*eclaType.Var),
		next:     nil,
		previous: s,
	}
}

func (s *Scope) GoUp() {
	cursor := s
	if cursor.next != nil {
		for cursor.next.next != nil {
			cursor = cursor.next
		}
	}
	cursor.next = nil
}

func (s *Scope) SetNextScope(next *Scope) {
	s.next = next
}

func (s *Scope) GetNextScope() *Scope {
	return s.next
}
