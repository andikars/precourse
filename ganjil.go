package main

import "fmt"

func main() {

	n := 100

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
