package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	// fmt.Println(colors) //map[green:#4bf745 red:#ff0000]

	// var colors map[string]string //Empty map to be filled later
	// fmt.Println(colors) //map[]

	// colors := make(map[string]string) //Create an empty map using make function
	// fmt.Println(colors)               //map[]
	// colors["white"] = "#ffffff"
	// fmt.Println(colors) //map[white:#ffffff]

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"
	// delete(colors, 10)
	// fmt.Println(colors) //map[]

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
	/*
		Hex code for  red  is  #ff0000
		Hex code for  green  is  #4bf745
		Hex code for  white  is  #ffffff
	*/
}
