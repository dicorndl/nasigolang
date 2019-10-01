package main

import (
	"log"
	"testing"
)

func TestTwoStackQueue(t *testing.T) {
	mockData := []int{10, 11, 12, 13}
	queue := NewQueue(10)
	// enqueue
	for i := 0; i < len(mockData); i++ {
		err := queue.Enqueue(mockData[i])
		if err != nil {
			log.Fatal(err)
			t.Fail()
		}
	}
	if queue.StackEnq.Len() != len(mockData) {
		log.Fatalf("stack enq len should be %d but %d", len(mockData), queue.StackEnq.Len())
		t.Fail()
	}
	if queue.StackDeq.Len() != 0 {
		log.Fatalf("stack deq len should be %d but %d", 0, queue.StackDeq.Len())
		t.Fail()
	}
	// dequeue
	value, err := queue.Dequeue()
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	if value != mockData[0] {
		log.Fatalf("front of the queue should be %d but %d", mockData[0], value)
		t.Fail()
	}
	if queue.StackEnq.Len() != 0 {
		log.Fatalf("after dequeue, stack enq len should be %d but %d", 0, queue.StackEnq.Len())
		t.Fail()
	}
	if queue.StackDeq.Len() != len(mockData)-1 {
		log.Fatalf("after dequeue, stack deq len should be %d but %d", len(mockData)-1, queue.StackDeq.Len())
		t.Fail()
	}
	// enqueue again
	err = queue.Enqueue(14)
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	if queue.StackEnq.Len() != 1 {
		log.Fatalf("after enqueue again, stack enq len should be %d but %d", 1, queue.StackEnq.Len())
		t.Fail()
	}
	if queue.StackDeq.Len() != len(mockData)-1 {
		log.Fatalf("after enqueue again, stack deq len still have been %d but %d", len(mockData)-1, queue.StackDeq.Len())
		t.Fail()
	}
	// dequeue again
	value, err = queue.Dequeue()
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	if value != mockData[1] {
		log.Fatalf("peek value should be %d but %d", mockData[1], value)
		t.Fail()
	}
	if queue.StackEnq.Len() != 1 {
		log.Fatalf("after dequeue again, stack enq len still have been %d but %d", 1, queue.StackEnq.Len())
		t.Fail()
	}
	if queue.StackDeq.Len() != len(mockData)-2 {
		log.Fatalf("after dequeue again, stack deq len still have been %d but %d", len(mockData)-2, queue.StackDeq.Len())
		t.Fail()
	}
	// peek
	value, err = queue.Peek()
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	if value != mockData[2] {
		log.Fatalf("peek value should be %d but %d", mockData[2], value)
		t.Fail()
	}
	if queue.StackEnq.Len() != 1 {
		log.Fatalf("after peek, stack enq len still have been %d but %d", 1, queue.StackEnq.Len())
		t.Fail()
	}
	if queue.StackDeq.Len() != len(mockData)-2 {
		log.Fatalf("after peek, stack deq len should not been changed %d but %d", len(mockData)-2, queue.StackDeq.Len())
		t.Fail()
	}
}