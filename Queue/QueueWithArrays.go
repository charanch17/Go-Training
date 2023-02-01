package main

import "fmt"

type Queue struct {
	rear  int
	queue [10000]int
}

func (q *Queue) setDefaults() {
	q.rear = -1
}

func (q *Queue) Enqueue(val int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}

	}()
	if q.rear == 9999 {
		panic("max queue size reached")
	} else {
		q.rear += 1
		q.queue[q.rear] = val
	}
}
func (q *Queue) Dequeue() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}

	}()
	if q.rear == -1 {
		panic("queue is empty")
	} else {
		var toRemove = q.queue[0]
		var temp [10000]int
		copy(temp[:], q.queue[1:])
		q.queue = temp
		q.rear -= 1
		return toRemove
	}

}
func (q *Queue) front() {
	fmt.Println("first element is", q.queue[0])
}
func (q *Queue) size() {
	fmt.Println(q.rear + 1)
}
func (q *Queue) display() {
	if q.rear == -1 {
		fmt.Println("queue is empty")

	} else {
		for i := 0; i <= q.rear; i++ {
			fmt.Printf("%v ", q.queue[i])
		}
	}
	fmt.Println("")
}
func main() {
	myqueue := Queue{}
	myqueue.setDefaults()
	myqueue.display()
	// trying to pop on empty stack
	myqueue.Dequeue()
	myqueue.size()
	myqueue.Enqueue(10)
	myqueue.Enqueue(20)
	myqueue.Enqueue(11)
	myqueue.Enqueue(21)
	myqueue.display()
	myqueue.front()
	myqueue.Dequeue()
	myqueue.display()
	myqueue.Dequeue()
	myqueue.display()
	myqueue.size()
	myqueue.display()
	myqueue.Dequeue()
	myqueue.Dequeue()
	myqueue.Dequeue()
	myqueue.Enqueue(11)
	myqueue.Enqueue(21)
	myqueue.display()
	myqueue.size()

}
