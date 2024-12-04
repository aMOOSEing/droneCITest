package main

import "fmt"

func Greeting(name string) string {
	return "Hello, " + name //testingg
}

func main() {
	input := Greeting("Mus")
	fmt.Println(input)
}
