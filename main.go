package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"main-mode/shape"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/scheduler"
)

// Выполняется самая первая
func init() {
	var initMsg string = "Init - выполняется быстрее main\n=======================\n"
	println(initMsg)
}

func main() {
	schedulerExample()
}

func interfaceExample() {
	var square = shape.NewSquare(50)
	var circle = shape.NewCircle(50)

	// Пустой интерфейс
	var emptyInterface = func(i interface{}) {
		// Проверка типа
		switch value := i.(type) {
		case int:
			fmt.Println("int", value)
		case string:
			fmt.Println("string", value)
		default:
			fmt.Println("unknown type", value)
		}

		// Проверка на существование типа (type guard)
		var str, ok = i.(string)
		if !ok {
			fmt.Println("interface is not a string")
			return
		}

		fmt.Println("interface: ", len(str))
	}

	// logs
	shape.PrintShapeArea(square)
	shape.PrintShapeArea(circle)

	emptyInterface(2222)
	emptyInterface("circle")

}

// Испольнование внешнего пакета
func schedulerExample() {
	var parseSubscriptionData = func(ctx context.Context) {
		time.Sleep(time.Second * 1)
		fmt.Printf("subscription parsed successfuly at %s\n", time.Now().String())
	}

	var sendStatistics = func(ctx context.Context) {
		time.Sleep(time.Second * 5)
		fmt.Printf("statistics sent at %s\n", time.Now().String())
	}

	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, parseSubscriptionData, time.Second*5)
	worker.Add(ctx, sendStatistics, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()

}

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

// simple HTTP_Server
func HTTP_Server() {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name,omitempty"`
	}

	var (
		users = []User{{1, "Tom"}, {2, "Bob"}}
	)

	var handleUsers = func(w http.ResponseWriter, r *http.Request) {
		resp, err := json.Marshal(users)
		fmt.Println("SERVER STARTED")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)

	}

	http.HandleFunc("/users", handleUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

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
