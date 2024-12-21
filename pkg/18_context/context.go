package context_pkg

import (
	"context"
	"fmt"
	"time"
)

func Main() {
	// Создаем контекст с таймаутом в 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Важно всегда вызывать cancel для освобождения ресурсов

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Операция превысила таймаут")
	case <-ctx.Done():
		fmt.Println("Контекст отменен:", ctx.Err())
	}

	// =====================
	// Создаем контекст со значением
	ctx = context.WithValue(context.Background(), "key", "value")

	// Получаем значение
	if value, ok := ctx.Value("key").(string); ok {
		fmt.Println("Получено значение:", value)
	}
}
