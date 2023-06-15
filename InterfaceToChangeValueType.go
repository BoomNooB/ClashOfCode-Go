//Normal way

// package main

// import "fmt"

// func main() {
// 	var N int
// 	fmt.Scan(&N)

// 	if N < 0 {
// 		fmt.Println("no negative numbers")
// 	} else if N%3 == 0 {
// 		fmt.Println(true)
// 	} else {
// 		fmt.Println(false)
// 	}
// }

// interface
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N < 0 {
		fmt.Println("no negative numbers")
	} else if N%3 == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
