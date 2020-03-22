package main

func makeWelcomer() func() {
	return makePrinter("Elo Ziemio!")
}

func makePrinter(data string) func() {
	sink, _ := makeMemorySink()
	return func() { sink(data) }
}
