package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	N int
	X []int
}

func NewInput() (*Input, error) {
	var N int
	var X []int

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		s := scanner.Text()
		fmt.Sscanf(s, "%d", &N)
	}

	X = make([]int, 0, N)

	if scanner.Scan() {
		s := scanner.Text()
		for _, item := range strings.Split(s, " ") {
			xi, err := strconv.Atoi(item)
			if err != nil {
				return nil, err
			}

			X = append(X, xi)
		}
	}

	return &Input{
		N: N,
		X: X,
	}, nil
}

func calc(P int, X []int) int {
	var total int

	for _, item := range X {
		sub := P - item
		energy := sub * sub
		total += energy
	}

	return total
}

func main() {
	input, err := NewInput()
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(input.X, func(i, j int) bool {
		return input.X[i] < input.X[j]
	})

	sub := input.X[input.N-1] - input.X[0]

	var min int
	for i := 0; i < sub; i++ {
		P := input.X[0] + i
		total := calc(P, input.X)

		if min == 0 || min > total {
			min = total
		}
	}

	fmt.Println(min)
}
