package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B int

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		s := scanner.Text()
		fmt.Sscanf(s, "%d %d", &A, &B)
	}

	curr := 1
	var count int
	for {
		if curr >= B {
			break
		}

		curr += A - 1

		count++
	}

	fmt.Println(count)
}
