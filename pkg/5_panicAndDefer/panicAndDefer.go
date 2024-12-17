package panicAndDefer

import "fmt"

// panic и defer
func panicAndDeferExample() {
	// panicHandler - обработчик паники (ошибки)
	var panicHandler = func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Возврат из recover")
		}

		fmt.Println("panicHandler() выполнен")
	}

	// panic - воспроизводит ошибку
	var panic = func() {
		var arr = []int{1, 2, 3}
		fmt.Println(arr[4])
		panic("panic example")
	} // panic - всплытие

	// defer - отложенное выполнение в конец
	defer panicHandler()
	panic()

	// этот код не выполнится, тк сработает паника
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

}
