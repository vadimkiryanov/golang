package channels

import "fmt"

func main() {
	// Создаем буферизованный канал с емкостью 2
	ch := make(chan int, 2)

	// Отправляем данные в канал
	ch <- 1
	ch <- 2

	// Принимаем данные из канала
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}
