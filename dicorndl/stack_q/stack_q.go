package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []int
	count int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() int {
	s.count--
	return s.nodes[s.count]
}

func Dequeue(s1, s2 *Stack) int {
	if s2.count == 0 {
		for s1.count > 0 {
			s2.Push(s1.Pop())
		}
	}

	return s2.Pop()
}

func main() {
	stack1 := NewStack()
	stack2 := NewStack()

	reader := bufio.NewReader(os.Stdin)
	inputQ, _ := reader.ReadString('\n')
	suffixTrimmedQ := strings.TrimSuffix(inputQ, "\n")
	q, _ := strconv.Atoi(suffixTrimmedQ)
	for i := 0; i < q; i++ {
		inputX, _ := reader.ReadString('\n')
		suffixTrimmedX := strings.TrimSuffix(inputX, "\n")
		fields := strings.Fields(suffixTrimmedX)
		if fields[0] == "1" {
			v, _ := strconv.Atoi(fields[1])
			stack1.Push(v)
		} else if fields[0] == "2" {
			Dequeue(stack1, stack2)
		} else if fields[0] == "3" {
			v := Dequeue(stack1, stack2)
			fmt.Println(v)
			stack2.Push(v)
		}
	}
}
