package main

import "fmt"

func makeWelcomer() func() {
	return makePrinter("Elo Ziemio!")
}

func makePrinter(data string) func() {
	return func() {
		fmt.Println(data)
	}
}
