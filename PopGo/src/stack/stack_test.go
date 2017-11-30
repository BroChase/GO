package stack

import "testing"

func TestStack(t *testing.T) {

    s := New();
	
	// Test 1
	if !s.IsEmpty() {
	    t.Errorf("Stack is not empty.")
	}
		
	// Test 2
	t2_v := 1
	if err := s.Push(t2_v); err != nil {
	    t.Errorf("Push error: %v", err.Error())
	}
	if s.Top() != t2_v {
	    t.Errorf("Top value = %v, want %v", s.Top, t2_v)
	}
	
	// Test 3
	var t3_v interface{}
	if err := s.Pop(&t3_v); err != nil {
	    t.Errorf("Pop error: %v", err.Error())
	} else if t3_v != t2_v {
	    t.Errorf("Pop value = %v, want %v", t3_v, t2_v)
	}
	
	// Test 4
	if !s.IsEmpty() {
	    t.Errorf("Stack is not empty.")
	}
	
	// Test 5
	var t5_v interface{}
	if err := s.Pop(&t5_v); err != Underflow {
	    t.Errorf("Pop did not return underflow")
	}
}