// Provides stack operations for any type. A stack is a first-in-first-out
// (LIFO) data store.
//
// This implementation provides an unbounded stack inplemented as a linked list
// of cells.
//
package stack

import "errors"

// Underflow is a error that occurs when you attempt to access as empty stack.
//
var Underflow = errors.New("stack underflow");

// A cell stores a single value of any type and a pointer to the next cell.
//
type cell struct {
    next *cell
	value interface{}
}

// A stack contains a pointer to the first cell.
//
type Stack struct {
    top *cell
}

func New() Stack {
    return Stack{nil}
}

// Pushes a value onto the stack.
//
func (s *Stack) Push(v interface{}) error {
    s.top = &cell{s.top, v}
	return nil
}

// Pops a value from the stack. An underflow error is returned if the stack is
// empty.
//
func (s *Stack) Pop(v *interface{}) error {
    if s.top == nil {
	    return Underflow
	}
	*v = s.top.value
	s.top = s.top.next
	return nil
}
func (s * Stack) PopOff() (interface{}, error){
	if s.top == nil{
		return nil, Underflow
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}
// Returns the top value on the stack. An underflow error is returned if the
// stack is empty.
//
func (s Stack) Top() interface{} {
    if s.top == nil {
	    return nil
	}
	return s.top.value
}

// Returns true is the stack is empty.
//
func (s Stack) IsEmpty() bool {
    return s.top == nil
}