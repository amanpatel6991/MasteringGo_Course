package main

import "fmt"

type Stack []int

func (s Stack) Empty() bool {
	return len(s) == 0
}
func (s *Stack) Push(v int) {
	(*s) = append((*s), v)
}
func (s *Stack) Pop() int {
	v := (*s)[len(*s) - 1]
	(*s) = (*s)[:len(*s) - 1]
	return v
}

type Queue []int

func (q Queue) Empty() bool {
	return len(q) == 0
}
func (q *Queue) Enqueue(v int) {
	(*q) = append((*q), v)
}
func (q *Queue) Dequeue() int {
	v := (*q)[0]
	(*q) = (*q)[1:len(*q)]
	return v
}

func main() {

	s := Stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
	// Output:
	// 2
	// 1
	// true

	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Empty())
	// Output:
	// 1
	// 2
	// true

}
