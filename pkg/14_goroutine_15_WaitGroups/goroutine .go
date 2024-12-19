package goroutine

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Определяем возможные действия для логов
var actions = []string{
	"logged IN",     // Пользователь вошел в систему
	"logged OUT",    // Пользователь вышел из системы
	"CREATE record", // Создание записи
	"DELETE record", // Удаление записи
	"UPDATE record", // Обновление записи
}

// Структура для хранения информации о действии и времени его выполнения
type logItem struct {
	action    string    // Действие, выполненное пользователем
	timestamp time.Time // Время выполнения действия
}

// Структура для хранения информации о пользователе
type user struct {
	id    int       // Уникальный идентификатор пользователя
	email string    // Электронная почта пользователя
	logs  []logItem // Список логов действий пользователя
}

// Пример пользователя с предопределенными логами
var userMock = user{
	id:    1,
	email: "some@mail.ru",
	logs: []logItem{
		{action: actions[0], timestamp: time.Now()},
		{action: actions[3], timestamp: time.Now()},
		{action: actions[2], timestamp: time.Now()},
		{action: actions[1], timestamp: time.Now()},
		{action: actions[0], timestamp: time.Now()},
		{action: actions[2], timestamp: time.Now()},
	},
}

// Метод для получения информации о действиях пользователя
func (u user) getActivityInfo() string {
	var out = fmt.Sprintf("ID:%d | Email:%s\n ActivityLog: \n", u.id, u.email)

	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.timestamp)
	}

	return out
}

// Функция для генерации и создания пользователей
func generateUsers(count int) []user {
	var users = make([]user, count)

	// Вложенная функция для генерации логов
	var generateLogs = func(actionIndex int) []logItem {
		var logs = make([]logItem, actionIndex)

		for i := 0; i < actionIndex; i++ {
			// Генерируем случайный индекс для выбора действия
			var lenActions = rand.Intn(len(actions))
			logs[i] = logItem{
				timestamp: time.Now(),
				action:    actions[lenActions],
			}
		}

		return logs
	}

	// Генерация пользователей
	for i := 0; i < count; i++ {
		users[i] = user{
			id:    i + 1,
			email: fmt.Sprintf("user%d@google.com", i+1),
			logs:  generateLogs(rand.Intn(100)), // Генерация случайного количества логов
		}
	}

	return users
}

// Функция для сохранения информации о пользователе в файл
func saveUserInfo(u user, wGoup *sync.WaitGroup) error {
	fmt.Printf("WRITING FILE FOR USER ID: %d\n", u.id)

	// Задержка дополнительная
	time.Sleep(time.Millisecond * 10)

	var fileName = fmt.Sprintf("logs/uid_%d.txt", u.id)

	// os.O_RDWR|os.O_CREATE - флаги для чтения и создания, 0644 - права доступа на запись и чтение
	var file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	var _, errWriteFile = file.WriteString(u.getActivityInfo())

	if errWriteFile != nil {
		return errWriteFile
	}

	wGoup.Done() // Уведомляем группу о том, что горутина завершилась

	return nil
}

// ===========================================================

// Функция для генерации пользователей и создания логов
func Main() {
	var timeCurrent = time.Now()      // Текущее время
	var waitGroup = &sync.WaitGroup{} // Группа ожидания горутин

	var users = generateUsers(500) // Сгенерированные пользователи

	for _, u := range users {
		waitGroup.Add(1)              // Добавляем горутину в ожидание завершения
		go saveUserInfo(u, waitGroup) // Запускаем горутину, в которой настроено завершение выполнеия горутины waitGroup.Done()
	}

	waitGroup.Wait() // Ожидаем завершения всех горутин, тем самым блокируя главную горутину

	fmt.Println("TIME ELAPSED: ", time.Since(timeCurrent).String())
}

// Пример использования горутин
func SleepRun() {
	go func() {
		fmt.Println("1")
	}()
	go fmt.Println("2")
	go fmt.Println("3")

	// Ожидание завершения горутин
	time.Sleep(time.Second)

	fmt.Println("FINISHED")
}
