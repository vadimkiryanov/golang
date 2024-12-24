package main

import wp_pkg "main-mode/pkg/19_wp"

// Выполняется самая первая
func init() {
	var initMsg string = "Init - выполняется быстрее main\n=======================\n"
	println(initMsg)
}

// Функция main - сама по себе как горутина
func main() {
	wp_pkg.Main()
}
