package main

import (
	context_pkg "main-mode/pkg/18_context"
)

// Выполняется самая первая
func init() {
	var initMsg string = "Init - выполняется быстрее main\n=======================\n"
	println(initMsg)
}

// Функция main - сама по себе как горутина
func main() {
	context_pkg.Main()
}
