package main

import "fmt"

func printEnHello() {
	fmt.Println("Hello, world")
}

func printRuHello() {
	fmt.Println("Привет, мир")
}

func printRuLove() {
	fmt.Println("Я люблю эту планету")
}

func printAnyRuHello(hello string) {
	fmt.Printf("Привет, %v\n", hello)
}

func printSizedStrings(firtsStr string, secondStr string) {
	fmt.Printf("%-70v %v\n", firtsStr, secondStr)
}

func main() {
	printEnHello()
	printRuHello()
	printRuLove()
	printAnyRuHello("Даша")
	printSizedStrings("Лупа", "Пупа")
	printSizedStrings("Кто такой Тайлер Дерден", "А ты не знаешь?")
}
