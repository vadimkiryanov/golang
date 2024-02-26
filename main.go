package main

import (
	"errors"
	"fmt"
)

func main() {

}

func closureIncrement() func() int {
	num := 0

	return func() int {
		num++
		fmt.Println("num = ", num)
		return num
	}

}

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

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 100 {
		return "Вход разрешен", nil

	} else if age < 18 {
		return "Вход запрещен", nil
	}

	return "Вход тоже запрещен", errors.New("do not enter the club")
}

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
