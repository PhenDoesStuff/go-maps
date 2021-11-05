package main

import "fmt"

func main() {
	// Below approach is taken if we want to figure out later what we want to add to it
	// var colors map[string]string

	// Pretty much equivalent to the code above
	colors := make(map[string]string)

	// Can add something to the map with the code below
	colors["white"] = "#ffffff"

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#008000",
	// }


	// Built in delete function to remove something from a map
	delete(colors, "white")

	fmt.Println(colors)
}