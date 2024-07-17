package main

import "fmt"

func main() {
	fmt.Println("2. В функции не передаются аргументы, но функция возвращает значение (значения)")
	fmt.Print(GetWelcomeMessage())
	fmt.Print(GetPiValue())
	fmt.Print(IsServerActive())
}

func GetWelcomeMessage() string {
	return "Welcome!"
}

func GetPiValue() float32 {
	return 3.14
}

func IsServerActive() bool {
	return true
}
