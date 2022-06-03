package main

import (
	"data-structures-in-go/Queue"
	"fmt"
	"math/rand"
)

const (
	N = 10
)

func main() {
	q := new(Queue.Queue)
	q.Create()

	for i := 0; i < N; i++ {
		go func() {
			q.Enqueue(new(Queue.Node).Create(rand.Int() % 10))
		}()
	}
	fmt.Println(q)
	i := q.Dequeue()
	for i != nil {
		fmt.Println(i)
		i = q.Dequeue()
	}
	fmt.Println(q)
}
