package main

import (
	"fmt"
	"math/rand"
	"time"
)

type callbackFunc func(stat BulkWriterStatus)

type BulkWriterStatus struct {
	Status int
	Datum  interface{}
}

type BulkWriterOperation struct {
	a BulkWriterStatus
	f callbackFunc
}

type BulkWriter struct {
	C      string
	Queue  []BulkWriterOperation
	numOps int
}

func (b *BulkWriter) enqueue(op BulkWriterOperation) {
	b.Queue = append(b.Queue, op)
}

func (b *BulkWriter) dequeue() *BulkWriterOperation {
	if len(b.Queue) <= 0 {
		return nil
	}

	op, q := b.Queue[0], b.Queue[1:]
	b.Queue = q
	return &op
}

func (b *BulkWriter) Do(datum interface{}, fn callbackFunc) {
	b.numOps++

	s := BulkWriterStatus{Datum: datum, Status: 0}

	b.enqueue(BulkWriterOperation{a: s, f: fn})
	go b.execute()

}

func (b *BulkWriter) execute() {
	op := b.dequeue()
	op.a.Status = fakeBlockingCall()
	op.f(op.a)
	b.numOps--
}

func (b *BulkWriter) Close() {
	for b.numOps > 0 {
	}
}

func fakeBlockingCall() int {
	rand.Seed(42)
	o := rand.Intn(10)
	t := rand.Int63n(2000)
	time.Sleep(time.Duration(t))
	return o
}

func main() {
	bw := BulkWriter{C: "my bulkwriter"}
	defer bw.Close()

	fn := func(s BulkWriterStatus) {
		fmt.Println("Callback!")
		fmt.Println(s.Datum)
		fmt.Println(s.Status)
	}

	bw.Do("foo", fn)
	bw.Do("bar", fn)
	bw.Do("baz", fn)

	fmt.Printf("%v\n", bw.Queue)
	for op := range bw.Queue {
		fmt.Println(op)
	}
	fmt.Printf("%v\n", bw.Queue)

}
