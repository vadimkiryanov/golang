package main

import (
	HTTPServer "main-mode/pkg/24_HTTPServer"
)

// Выполняется самая первая
func init() {
	var initMsg string = "Init - выполняется быстрее main\n=======================\n"
	println(initMsg)
}

// Функция main - сама по себе как горутина
func main() {
	HTTPServer.Main()
}
