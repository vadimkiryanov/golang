package main

import (
	"errors"
	"fmt"
)

// Выполняется самая первая
var initMsg string = "Init - выполняется быстрее main\n=======================\n"

func init() {
	println(initMsg)
}

func main() {
	arrayAndSlice()
}

// массивы
func arrayAndSlice() {
	msgArr := [3]string{"a", "r", "r"}                     // У массива нельзя менять длину
	msgSlice := []string{"s", "l", "i", "c", "e"}          // У слайса можно менять длину и передается по ссылке
	msgSliceByMake := make([]string, 5)                    // Создание слайса через make
	msgSliceByMakeAppended := append(msgSliceByMake, "22") // Если не хватает длины при добавлении, то длина увеличивается на x2

	var changeSlice = func(msgSlice []string) []string {
		msgSlice[0] = "22"
		return msgSlice
	}

	fmt.Println(msgArr, " | msgArr")
	fmt.Println(msgSlice, " | msgSlice")
	fmt.Println(changeSlice(msgSlice), " | changeSlice")

	fmt.Println(msgSliceByMake, " | msgSliceByMake", len(msgSliceByMake), cap(msgSliceByMake))
	fmt.Println(msgSliceByMakeAppended, " | msgSliceByMakeAppended", len(msgSliceByMakeAppended), cap(msgSliceByMakeAppended))

	// matrix - матрица
	fmt.Println("matrix\n=======================\n'")
	matrix := make([][]int, 10)

	var counter int = 0
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)

		for j := 0; j < 10; j++ {
			counter++
			matrix[i][j] = counter
		}

		fmt.Println(matrix[i])

	}

	// range - цикл
	fmt.Println("range\n=======================\n'")
	for index, value := range matrix {
    if index == 6 {
      fmt.Println("index равен: ", 6, " остановка цикла")
      break
    }
		fmt.Println(index, value)
	}

}

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

// Замыкание, Анонимная функция
func closureIncrement() func() int {
	num := 0

	return func() int {
		num++
		fmt.Println("num = ", num)
		return num
	}

}

// Неограниченные аргументы
func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return min
}

func sayHello(name string, age int) string {

	message := fmt.Sprintf("Привет, %s! Тебе %d лет!", name, age)

	return message
}

// Функции Условныи оператор if Обработка ошибок
func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 100 {
		return "Вход разрешен", nil

	} else if age < 18 {
		return "Вход запрещен", nil
	}

	return "Вход тоже запрещен", errors.New("do not enter the club")
}

// Конструкция switch case
func predication(dayOfTheWeek string) (string, error) {
	switch dayOfTheWeek {
	case "пн":
		return "Сегодня понедельник", nil
	case "вт":
		return "Сегодня вторник", nil
	case "ср":
		return "Сегодня среда", nil
	case "чт":
		return "Сегодня четверг", nil
	case "пт":
		return "Сегодня пятница", nil
	case "сб":
		return "Сегодня суббота", nil
	case "вс":
		return "Сегодня воскресенье", nil

	default:
		return "Такого дня недели нет", errors.New("not excepted day of the week")
	}

}
