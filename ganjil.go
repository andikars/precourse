package main

import "fmt"

func main() {
	ganjil(1, 100)
}

func ganjil(awal uint, akhir uint) {
	if awal < akhir && awal > 0 {
		for i := awal; i <= akhir; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}
}
