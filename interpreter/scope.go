package interpreter

import "github.com/tot0p/Ecla/interpreter/eclaType"

type ScopeType int

const (
	SCOPE_MAIN ScopeType = iota
	SCOPE_FUNCTION
	SCOPE_LOOP
	SCOPE_CONDITION
	SCOPE_TRY
	SCOPE_CATCH
	SCOPE_FINALLY
)

type Scope struct {
	Var      map[string]*eclaType.Var
	next     *Scope
	previous *Scope
	Type     ScopeType
	InFunc   bool
}

func NewScopeMain() *Scope {
	return &Scope{
		Var:      make(map[string]*eclaType.Var),
		next:     nil,
		previous: nil,
		Type:     SCOPE_MAIN,
		InFunc:   false,
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

func (s *Scope) GoDeep(Type ScopeType) {
	cursor := s
	for cursor.next != nil {
		cursor = cursor.next
	}
	InFunc := cursor.InFunc
	if Type == SCOPE_FUNCTION {
		InFunc = true
	}
	cursor.next = &Scope{
		Var:      make(map[string]*eclaType.Var),
		next:     nil,
		previous: cursor,
		Type:     Type,
		InFunc:   InFunc,
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

func (s *Scope) GetFunctionScope() *Scope {
	cursor := s
	for cursor.Type != SCOPE_FUNCTION && cursor.previous != nil {
		cursor = cursor.previous
	}
	if cursor.Type != SCOPE_FUNCTION {
		return nil
	}
	return cursor
}

func (s *Scope) InFunction() bool {
	return s.InFunc
}
