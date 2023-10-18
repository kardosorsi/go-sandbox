package main

import "fmt"

type bot interface { 
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type dutchBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	db := dutchBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(db)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic for englishBot
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// custom logic for spanishBot
	return "Hola!"
}

func (dutchBot) getGreeting() string {
	// custom logic for dutchBot
	return "Hoi!"
}