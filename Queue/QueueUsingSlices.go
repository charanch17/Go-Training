package main

import "fmt"

type Queue struct {
	queue []int
}

func (q *Queue) Enqueue(val int) {
	q.queue = append(q.queue, val)
}

func (q *Queue) Dequeue() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if len(q.queue) > 0 {
		toRemove := q.queue[0]
		q.queue = q.queue[1:]
		return toRemove
	} else {
		panic("queue is empty")
	}
}
func (q *Queue) display() {
	for _, i := range q.queue {
		fmt.Printf("%v ", i)
	}
	fmt.Println("")
}
func (q *Queue) size() int {
	fmt.Println("size", len(q.queue))
	return len(q.queue)
}
func (q *Queue) front() {
	if len(q.queue) > 0 {
		fmt.Println("first element is ", q.queue[0])
	} else {
		fmt.Println("queue is empty")
	}
}

func main() {
	myqueue := Queue{}
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
