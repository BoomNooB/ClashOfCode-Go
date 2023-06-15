package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	var M int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &M)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	var totalNumber int64

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	for i := 0; i < N; i++ {
		E, _ := strconv.ParseInt(inputs[i], 10, 32)
		_ = E

		totalNumber = totalNumber + (E % int64(M))
	}

	fmt.Println(totalNumber)

}
