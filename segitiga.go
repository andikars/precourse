package main

import "fmt"

func main() {
	segitiga(4)
}

func segitiga(n int) {
	if n > 0 {
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("*")
			}
			fmt.Println("")
		}
	}
}
