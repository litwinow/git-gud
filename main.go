package main

import "fmt"

func makeWelcomer() func() {
	return func() {
		fmt.Println("Elo Ziemio!")
	}
}

func main() {
	welcomer := makeWelcomer()
	welcomer()
}
