package main

import (
	"fmt"
	"os"
)

type Sink func(string)

func makeStdoutSink() Sink {
	return func(data string) {
		fmt.Println(data)
	}
}

func makeFileSink(filename string) Sink {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return func(data string) {
		file.WriteString(data)
	}
}

func makeMemorySink() (Sink, []string) {
	buf := make([]string, 10)
	currentPos := 0
	return func(data string) {
		buf[currentPos] = data
		currentPos++
		if currentPos > 10 {
			panic("memory sink capacity reached")
		}
	}, buf
}
