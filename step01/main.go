package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func extractNums(tokens []string, i int) int {
	b, _ := strconv.Atoi(tokens[i + 2])
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()

		tokens := strings.Fields(text)
		result, _ := strconv.Atoi(tokens[0])
		for i, v := range tokens[1:len(tokens) - 1] {
			b := extractNums(tokens, i)
			switch v {
			case "+":
				result = sum(result, b)
			case "-":
				result = sub(result, b)
			case "*":
				result = mul(result, b)
			case "/":
				result = div(result, b)
			default:
				continue
			}
		}

		fmt.Println(result)
	}
}
