package structPkg

import (
	"fmt"
	"main-mode/pkg/shape"
)

// Структуры
func structExample() {

	// Конструктор
	var NewUser = func(name string, age int, gender string, city string) shape.User {
		return shape.User{
			Name:   name,
			Age:    shape.Age(age),
			Gender: gender,
			City:   city,
		}
	}

	// Создание переменной вручную по структуре
	var user1 = shape.User{
		Name:   "Tom",
		Age:    12,
		Gender: "male",
		City:   "New York",
	}

	// Создание переменной генератором по структуре
	var user2 = NewUser("Bob", 23, "male", "New York")
	user1.GetName()                            // Вызов метода "printUserInfo", который был добавлен
	user2.SetName("НОВОЕ ИМЯ - ПОМЕНЯЛОСЬ!!!") // Вызов метода "printUserInfo", который был добавлен

	fmt.Println(user1)
	fmt.Println(user2) // Здесь поменялось значение
	fmt.Println(user1.Age.IsAdult())
}
