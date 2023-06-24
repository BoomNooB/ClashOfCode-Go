package main

import "fmt"

func main() {
	var str string
	twoCharString := func() string {
		if len(str) >= 2 {
			return str[:2]
		}
		return str
	}()

	fmt.Println(twoCharString)
}
