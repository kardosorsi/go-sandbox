package main

import "fmt"

func main() {
	oneToTen := []int{1,2,3,4,5,6,7,8,9,10}

	for _, val := range oneToTen {
		if val % 2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}