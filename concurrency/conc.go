package main

import (
	"fmt"
	"math/rand"
	"time"
)

type callbackFunc func(stat QueueStatus)

type QueueStatus struct {
	Status int
	Datum  interface{}
}

type QueueOperation struct {
	a QueueStatus
	f callbackFunc
}

type Queue struct {
	C      string
	Queue  []QueueOperation
	numOps int
}

func (b *Queue) enqueue(op QueueOperation) {
	b.Queue = append(b.Queue, op)
}

func (b *Queue) dequeue() *QueueOperation {
	if len(b.Queue) <= 0 {
		return nil
	}

	op, q := b.Queue[0], b.Queue[1:]
	b.Queue = q
	return &op
}

func (b *Queue) Do(datum interface{}, fn callbackFunc) {
	b.numOps++

	s := QueueStatus{Datum: datum, Status: 0}

	b.enqueue(QueueOperation{a: s, f: fn})
	go b.execute()

}

func (b *Queue) execute() {
	op := b.dequeue()
	op.a.Status = fakeBlockingCall()
	op.f(op.a)
	b.numOps--
}

func (b *Queue) Close() {
	for b.numOps > 0 {
	}
}

func fakeBlockingCall() int {
	rand.NewSource(time.Now().UnixNano())
	o := rand.Intn(10)
	t := rand.Int63n(2000)
	time.Sleep(time.Duration(t))
	return o
}

func main() {
	var q Queue
	defer q.Close()

	fn := func(s QueueStatus) {
		fmt.Println("Callback!")
		fmt.Println(s.Datum)
		fmt.Println(s.Status)
	}

	q.Do("one", fn)
	q.Do("two", fn)
	q.Do("three", fn)

	fmt.Printf("%v\n", q.Queue)
	for op := range q.Queue {
		fmt.Println(op)
	}
	fmt.Printf("%v\n", q.Queue)

}
