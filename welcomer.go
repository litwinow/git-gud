package main

func makeWelcomer() func() {
	return func() {
		fmt.Println("Elo Ziemio!")
	}
}