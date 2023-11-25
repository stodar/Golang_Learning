package main

import "fmt"

func main() {

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")
	// var colors map[string]string

	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#008000",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("the hex of color %+v is : %+v \n", color, hex)
	}
}
