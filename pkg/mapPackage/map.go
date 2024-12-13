package mapPackage

import "fmt"

// Мапы
func mapExample() {
	var hashMap = map[string]int{
		"a": 1,
		"b": 2,
	}

	var hashMapMake = make(map[string]int) // make позвояет инициализровать мапу
	hashMapMake["a"] = 1                   // добавление в пустую мапу возможно благодаря make, иначе была бы ошибка
	fmt.Println(hashMapMake)

	// добавление
	hashMap["c"] = 3
	// удаление
	delete(hashMap, "a")
	// exists - проверка на существование
	var value, exists = hashMap["a"]
	if exists {
		fmt.Println(value)
		fmt.Println(exists)
	} else {
		fmt.Println(hashMap)
	}

	// итерация
	for key, value := range hashMap {
		fmt.Println(`key:`, key, `value:`, value)
	}

	// Длина
	fmt.Println(len(hashMap))

}

