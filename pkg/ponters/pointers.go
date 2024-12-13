package ponters

import "fmt"

// Указатели, массивы и указатели цикл for | *string - указатель
func pointersExample() {
	// * - принятие
	// & - отдача
	var msg = "Hello world"
	println(msg, " | msg до 'changeMsg'")
	println(&msg, " Область памяти")

	var changeMsg func(*string) string = func(msg *string) string {
		*msg += " (pointers(msg))" // дереференс
		return *msg
	}

	fmt.Println(changeMsg(&msg) + " | msg после 'changeMsg'") // &msg - референс
}
