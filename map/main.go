package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// fmt.Println(colors)
	printMap(colors)

	var anotherColors map[string]string
	fmt.Println(anotherColors)

	yetAnotherColors := make(map[int]string)
	yetAnotherColors[10] = "#ffffff"
	delete(yetAnotherColors, 10)
	fmt.Println(yetAnotherColors)
}

func printMap(c map[string]string) { 
	for color, hex := range c {
		fmt.Println(color + ":" + hex)
	}
}