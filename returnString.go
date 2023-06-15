package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	scanner.Scan()
	inputString := scanner.Text()
	_ = inputString // to avoid unused error

	emptyString := "empty"

	if N == 0 {
		fmt.Println(emptyString)
		os.Exit(0)
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%v", inputString)
	}
}
