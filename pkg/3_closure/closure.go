package closure

import "fmt"

// Замыкание, Анонимная функция
func ClosureIncrement() func() int {
	num := 0

	return func() int {
		num++
		fmt.Println("num = ", num)
		return num
	}

}
