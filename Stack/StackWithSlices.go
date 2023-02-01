package main

import "fmt"

type Stack struct {
	stack []int
}

func (s *Stack) push(val int) {
	s.stack = append(s.stack, val)
}
func (s *Stack) pop() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if len(s.stack) == 0 {
		panic("unable to pop as stack is empty")
	} else {
		toRemove := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return toRemove
	}
}

func (s *Stack) size() {
	fmt.Println("size", len(s.stack))
}

func (s *Stack) display() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if len(s.stack) == 0 {
		panic("stack is empty")
	} else {
		for _, i := range s.stack {
			fmt.Printf("%v ", i)
		}
		fmt.Println("")
	}
}
func (s *Stack) top() {
	fmt.Println("top is ", s.stack[len(s.stack)-1])
}
func main() {

	mystack := Stack{}
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
