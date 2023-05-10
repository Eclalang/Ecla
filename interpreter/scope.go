package interpreter

import "github.com/Eclalang/Ecla/interpreter/eclaType"

// ScopeType is the type of scope.
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

// Scope is a scope.
type Scope struct {
	Var      map[string]*eclaType.Var
	next     *Scope
	previous *Scope
	Type     ScopeType
	InFunc   bool
}

// NewScopeMain returns a new main scope.
func NewScopeMain() *Scope {
	return &Scope{
		Var:      make(map[string]*eclaType.Var),
		next:     nil,
		previous: nil,
		Type:     SCOPE_MAIN,
		InFunc:   false,
	}
}

// Set sets the value of the variable with the given name.
func (s *Scope) Set(name string, value *eclaType.Var) {
	cursor := s
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.Var[name] = value
}

// Get returns the value of the variable with the given name.
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

// GoDeep creates a new scope was is deeper than the current one.
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

// GoUp goes up in the scope and deletes the current one.
func (s *Scope) GoUp() {
	cursor := s
	if cursor.next != nil {
		for cursor.next.next != nil {
			cursor = cursor.next
		}
	}
	cursor.next = nil
}

// SetNextScope sets the next scope.
func (s *Scope) SetNextScope(next *Scope) {
	s.next = next
}

// GetNextScope returns the next scope.
func (s *Scope) GetNextScope() *Scope {
	return s.next
}

// GetFunctionScope returns the function scope.
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

// InFunction returns true if the scope is in a function.
func (s *Scope) InFunction() bool {
	return s.InFunc
}
