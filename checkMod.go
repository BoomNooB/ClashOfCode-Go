package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var m int
	fmt.Scan(&m) // m คือจำนวนเลขในลูป

	for i := 0; i < m; i++ {
		var n int
		fmt.Scan(&n)
		if n%2 == 0 {
			fmt.Println(n * 2)
		} else {
			fmt.Println(n * 3)
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
}
