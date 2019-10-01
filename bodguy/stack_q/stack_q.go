package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	// ErrStackOverflow stack overflow error
	ErrStackOverflow = errors.New("stack overflow")
	// ErrStackUnderflow stack underflow error
	ErrStackUnderflow = errors.New("stack underflow")
	// ErrTooLargeNumber too large number error
	ErrTooLargeNumber = errors.New("1 <= |x| <= 10^9 error")
	// ErrTooManyCommand too many command error
	ErrTooManyCommand = errors.New("1 <= q <= 10^5 error")
	// ErrUnknownCommand unknown command error
	ErrUnknownCommand = errors.New("1 <= type <= 3 error")
)

// Stack ADT
type Stack struct {
	top   int
	value []int
}

// NewStack stack ADT contstructor
func NewStack(len int) *Stack {
	return &Stack{
		top:   -1,
		value: make([]int, len),
	}
}

// Push push a single value to top of the stack
func (s *Stack) Push(value int) error {
	if s.top+1 >= len(s.value) {
		return ErrStackOverflow
	}
	s.top++
	s.value[s.top] = value
	return nil
}

// Pop pop a single value from the top of the stack
func (s *Stack) Pop() (int, error) {
	if s.top < 0 {
		return 0, ErrStackUnderflow
	}
	value := s.value[s.top]
	s.top = s.top - 1
	return value, nil
}

// Len retreive stack length
func (s *Stack) Len() int {
	return s.top + 1
}

// Peek just peek from stack, not popping
func (s *Stack) Peek() (int, error) {
	if s.top < 0 {
		return 0, ErrStackUnderflow
	}
	return s.value[s.top], nil
}

// Queue ADT compose of two stacks
type Queue struct {
	StackEnq *Stack
	StackDeq *Stack
}

// NewQueue queue ADT constructor
func NewQueue(len int) *Queue {
	return &Queue{
		NewStack(len),
		NewStack(len),
	}
}

// Enqueue enqueue a single value to back of the queue
func (q *Queue) Enqueue(value int) error {
	if value < 1 || value > int(math.Pow10(9)) {
		return ErrTooLargeNumber
	}
	return q.StackEnq.Push(value)
}

// Dequeue dequeue a single value from front of the queue
func (q *Queue) Dequeue() (int, error) {
	err := q.move()
	if err != nil {
		return 0, err
	}
	return q.StackDeq.Pop()
}

// Peek just peek a value from front of queue
func (q *Queue) Peek() (int, error) {
	err := q.move()
	if err != nil {
		return 0, err
	}
	return q.StackDeq.Peek()
}

func (q *Queue) move() error {
	if q.StackDeq.Len() == 0 {
		len := q.StackEnq.Len()
		for i := 0; i < len; i++ {
			value, err := q.StackEnq.Pop()
			if err != nil {
				return err
			}
			err = q.StackDeq.Push(value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Panic dispose all kind of error to panic!
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// scan from terminal input
	scanner := bufio.NewScanner(os.Stdin)
	// first command should be a number of command
	var stackSize int
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		Panic(err)
		if num < 1 || num > int(math.Pow10(5)) {
			panic(ErrTooManyCommand)
		}
		stackSize = num
	}

	queue := NewQueue(stackSize)

	for i := 0; i < stackSize; i++ {
		if scanner.Scan() {
			text := scanner.Text()
			token := strings.Fields(text)
			switch token[0] {
			case "1": // enqueue
				value, err := strconv.Atoi(token[1])
				Panic(err)
				err = queue.Enqueue(value)
				Panic(err)
			case "2": // dequeue
				_, err := queue.Dequeue()
				Panic(err)
			case "3": // print at the very front element
				value, err := queue.Peek()
				Panic(err)
				fmt.Printf("%d\n", value)
			default:
				panic(ErrUnknownCommand)
			}
		}
	}
}