package wp_pkg

import (
	"fmt"
	"time"
)

// Main демонстрирует работу с горутинами и каналами для параллельной обработки задач
func Main() {
    // Засекаем время начала выполнения
    var tNow = time.Now()
    
    // Определяем количество задач и работников
    var jobsCount, workerCount = 15, 10

    // Создаем буферизированные каналы для задач и результатов
    // Размер буфера равен количеству задач
    var jobs = make(chan int, jobsCount)
    var results = make(chan int, jobsCount)

    // Запускаем workerCount горутин-работников
    for i := 0; i < workerCount; i++ {
        go worker(i+1, jobs, results)
    }

    // Отправляем задачи в канал jobs
    for i := 0; i < jobsCount; i++ {
        jobs <- i + 1
    }
    // Закрываем канал jobs, сигнализируя что новых задач не будет
    close(jobs)

    // Получаем и выводим результаты выполнения всех задач
    for i := 0; i < jobsCount; i++ {
        fmt.Printf("result #%d : value = %d\n", i+1, <-results)
    }

    // Выводим общее время выполнения всех задач
    fmt.Printf("time ELAPSED = %s\n", time.Since(tNow).String())
}

// worker обрабатывает задачи из канала jobs и отправляет результаты в канал results
// Параметры:
//   - id: уникальный идентификатор работника
//   - jobs: канал только для чтения задач
//   - results: канал только для записи результатов
func worker(id int, jobs <-chan int, results chan<- int) {
    // Читаем задачи из канала jobs пока он не закрыт
    for job := range jobs {
        // Имитируем работу паузой в 1 секунду
        time.Sleep(time.Second)
        fmt.Printf("worker %d finished\n", id)
        // Отправляем квадрат числа в канал результатов
        results <- job * job
    }
}