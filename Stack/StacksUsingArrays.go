package main

import (
	"fmt"
)

type Stack struct {
	stack   [10000]int
	current int
}

// pushes element at the end of the stack
func (s *Stack) push(val int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}

	}()
	if s.current < 9999 {
		s.current += 1
		s.stack[s.current] = val
	} else {
		panic("Stack overflow")
	}

}
func (s *Stack) pop() int {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}

	}()
	if s.current == -1 {
		panic("unable to pop as stack is empty")
	}
	toRemove := s.stack[s.current]
	s.current -= 1
	return toRemove
}

func (s *Stack) display() {
	if s.current == -1 {
		fmt.Println("Stack is empty")
	}
	for i := 0; i <= s.current; i++ {
		fmt.Printf("%v ", s.stack[i])
	}
	fmt.Println("")
}

func (s *Stack) setDefaults() {
	s.current = -1

}
func (s *Stack) top() {
	if s.current > -1 {
		fmt.Println("top is ", s.stack[s.current])
	} else {
		fmt.Println("stack is empty")
	}
}

func (s *Stack) size() {
	fmt.Println("stacksize is", s.current+1)
}

func main() {

	mystack := Stack{}
	mystack.setDefaults()
	mystack.display()
	// trying to pop on empty stack
	mystack.pop()
	mystack.size()
	mystack.push(10)
	mystack.push(20)
	mystack.push(11)
	mystack.push(21)
	mystack.display()
	mystack.top()
	mystack.pop()
	mystack.pop()
	mystack.size()
	mystack.display()
	mystack.pop()
	mystack.pop()
	mystack.pop()
	mystack.push(11)
	mystack.push(21)
	mystack.display()
	mystack.size()

}
