package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englistBot struct{}
type spanishBot struct{}

func main() {
	eb := englistBot{}
	sb := spanishBot{}

	// printEnglishGreeting(eb) //Hello There!
	// printSpanishGreeting(sb) //Hola!

	printGreeting(eb) //Hello There!
	printGreeting(sb) //Hola!
}

// func printEnglishGreeting(eb englistBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printSpanishGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englistBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello There!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}
