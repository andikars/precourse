package main

import "fmt"

func main() {

	n := 50

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("buzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}

}
