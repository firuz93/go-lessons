package main

import "fmt"

func main() {
	fmt.Println("3. В функции передаются аргументы, но функция ничего не возвращает")
	fmt.Print(PrintNumber(10))
	fmt.Print(GreetUser("Firuz"))
	fmt.Print(ToggleLight(false))
}

func PrintNumber(int Num) {
	return Num
}

func GreetUser(string str) {
	return "Hello " + str
}

//func ToggleLight(bool light) {
//  if light = true{
//	return "Light on"
//	else {
//		return "Light off"
//	}
//}
//return true
//}
