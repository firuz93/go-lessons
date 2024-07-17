package main

import "fmt"

func main() {
	fmt.Println("1. В функции не передаются аргументы и ничего не возвращает функция")
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()
}

func PrintGreeting() {
	fmt.Println("Hello world!")
}

func DisplaySeparator() {
	fmt.Println("--------------------")
}

func NotifyStart() {
	fmt.Println("Program started")
}
