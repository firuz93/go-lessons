package main

import "fmt"

func main() {
	// 1 Дана сторона квадрата a. Найти его периметр P = 4·a.
	var a = 4
	var p = 4 * a
	fmt.Println("zadacha 1")
	fmt.Println(p)
	//2 Дана сторона квадрата a. Найти его площадь S = a*a
	var s = a * a
	fmt.Println("zadacha 2")
	fmt.Println(s)
	//3 Даны стороны прямоугольника a и b. Найти его площадь S = a·b и периметр P = 2·(a + b).
	var b = 6
	var s3 = a * b
	var p3 = 2 * (a + b)
	fmt.Println("zadacha 3")
	fmt.Println(s3)
	fmt.Println(p3)
	//4 Дан диаметр окружности d. Найти ее длину L = π·d. В качестве значения π использовать 3.14.
	var d = 12.5
	var pi = 3.14
	var l = pi * d
	fmt.Println("zadacha 4")
	fmt.Println(l)
	//5 Дана длина ребра куба a. Найти объем куба V = a3 и площадь его поверхности S = 6·a
	var v = a * a * a
	var s5 = 6 * a
	fmt.Println("zadacha 5")
	fmt.Println("V =", v)
	fmt.Println("S =", s5)
	//6 Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти его объем V = a·b·c и площадь поверхности S = 2·(a·b + b·c + a·c).
	var a6 = 2
	var b6 = 3
	var c6 = 4
	var v6 = a6 * b * c6
	var s6 = 2 * (a6*b6 + b6*c6 + a6*c6)
	fmt.Println("zadacha 6")
	fmt.Println("V =", v6)
	fmt.Println("S =", s6)
	//7 Найти длину окружности L и площадь круга S заданного радиуса R: L = 2·π·R, S = π·R
	//В качестве значения π использовать 3.14.
	var R = 5.75
	var l7 = 2 * 3.14 * R
	var s7 = 3.14 * R
	fmt.Println("zadacha 7")
	fmt.Println("l=", l7)
	fmt.Println("s=", s7)
	//8 Даны два числа a и b. Найти их среднее арифметическое: (a + b)/2.
	var a8 = 11
	var b8 = 13
	var srz = (a8 + b8) / 2
	fmt.Println("zadacha 8")
	fmt.Println("srz=", srz)
	//15 Даны целые положительные числа A и B (A > B). На отрезке длины A размещено максимально возможное количество отрезков длины B (без наложений).
	//Используя операцию взятия остатка от деления нацело, найти длину незанятой части отрезка A.
	var a15 = 184
	var b15 = 3
	var cd = a15 / b15
	var ost = a15 - cd*b15
	fmt.Println("zadacha 15")
	fmt.Println("остаток=", ost)

}
