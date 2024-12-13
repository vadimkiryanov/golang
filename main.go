package main

import (
	"main-mode/pkg/goroutine"
)

// Выполняется самая первая
func init() {
	var initMsg string = "Init - выполняется быстрее main\n=======================\n"
	println(initMsg)
}

// Функция main - сама по себе как горутина
func main() {
	goroutine.SleepRun()
}
