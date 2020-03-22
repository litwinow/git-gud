package main

func makeWelcomer() func() {
	sink, _ := makeMemorySink()
	return func() { sink("Elo Ziemio!") }
}
