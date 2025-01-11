package main

import sql_pkg "main-mode/pkg/25_SQL"

// Выполняется самая первая
func init() {
	var initMsg string = "Init - выполняется быстрее main\n=======================\n"
	println(initMsg)
}

// Функция main - сама по себе как горутина
func main() {
	sql_pkg.SQL() 
}
