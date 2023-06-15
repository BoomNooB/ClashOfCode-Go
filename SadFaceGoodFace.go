package main

import (
	"fmt"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Buffer(make([]byte, 1000000), 1000000)

	// scanner.Scan()
	// facesInput := scanner.Text()
	facesInput := "=) ^_^"

	// faces := strings.Split(facesInput, " ") // split into slice of string
	faces := strings.Fields(facesInput)

	fmt.Println("faces :", faces)
	fmt.Printf("%T", faces) // slice of string

	storePoint := 0

	for _, face := range faces {
		if face == "=)" || face == "^_^" {
			storePoint++
		} else if face == ":(" || face == ">_<" {
			storePoint--
		}
	}

	fmt.Println(storePoint) // Write answer to stdout
}
