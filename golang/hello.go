package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	power, powValue := getPower("400", 2000)
	fmt.Printf("%s %d\n", power, powValue)

	var name string
	name = "frank"
	fmt.Println(name)

	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
}

func getPower(a string, b int) (string, int) {
	if a == "400" || a == "200" {
		return a, b
	} else {
		return "default", 0
	}
}
