package main

import "fmt"

func main() {
	//For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//While
	i2 := 0
	for i2 < 10 {
		println("While")
		fmt.Println(i2)
		i2++
		if i2 > 10 {
			break
		}
	}
}
