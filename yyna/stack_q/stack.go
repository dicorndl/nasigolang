package main

import "fmt"

// https://www.hackerrank.com/challenges/queue-using-two-stacks/problem

type Stack struct {
	elements []int
	count    int
}

var inStack = newStack()
var outStack = newStack()

func main() {
	var count int
	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		var command, value int
		fmt.Scanln(&command, &value)
		
		switch {
		case command == 1:
			enQueue(value)
		case command == 2:
			deQueue()
		case command == 3:
			fmt.Println(peekQueue())
		}
	}
}

func newStack() *Stack {
	stack := Stack{}
	stack.count = 0
	stack.elements = []int{}
	return &stack
}

func push(stack *Stack, element int) {
	stack.count++
	stack.elements = append(stack.elements, element)
}

func pop(stack *Stack) int {
	var value int
	if stack.count > 0 {
		stack.count--
		value = stack.elements[stack.count]
		stack.elements = stack.elements[:stack.count]
	}
	return value

}

func enQueue(element int) {
	push(inStack, element)
}

func deQueue() {
	if outStack.count == 0 {
		for inStack.count != 0 {
			push(outStack, pop(inStack))
		}
	}

	pop(outStack)
}

func peekQueue() int {
	if outStack.count == 0 {
		for inStack.count != 0 {
			push(outStack, pop(inStack))
		}
	}

	return peek(outStack)
}

func peek(stack *Stack) int {
	if stack.count > 0 {
		return stack.elements[stack.count-1]
	}
	return -1
}
