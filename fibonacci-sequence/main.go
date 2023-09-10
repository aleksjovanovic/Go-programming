package main

import "fmt"

var sequence []int

func fS(n int) {

	for i := 1; i <= n; i++ {

		if i == 1 {
			sequence = append(sequence, 0)
		}

		if i == 2 {
			sequence = append(sequence, 1)
		}

		if i >= 3 {

			sequence = append(sequence, (sequence[(len(sequence)-2)] + sequence[(len(sequence)-1)]))

		}
		fmt.Println(sequence) // printing step by step
	}
	// fmt.Println(sequence) printing whole sequence
}

func main() {
	fS(20)
}
