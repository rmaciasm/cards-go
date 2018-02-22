package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "red",
		"verde": "green",
	}

	colors["morado"] = "purple"
	// delete(colors, "yellow")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for espaniol, ingles := range c {
		fmt.Println("Spanish for", ingles, " is ", espaniol)
	}
}
