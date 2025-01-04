package coincap

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// LoggingRoundTripper структура для логирования HTTP-запросов
// Реализует интерфейс http.RoundTripper для перехвата и логирования запросов
type LoggingRoundTripper struct {
	logger io.Writer         // интерфейс для записи логов
	next   http.RoundTripper // следующий обработчик в цепочке
}

// RoundTrip перехватывает HTTP-запрос, логирует его и передает дальше
// req: перехваченный HTTP-запрос
// возвращает ответ и ошибку от следующего обработчика
func (loggerRT LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Логируем время, метод и URL запроса
	fmt.Fprintf(loggerRT.logger,
		"=== loggingRoundTripper: [%s] %s %s \n",
		time.Now().Format(time.ANSIC), // текущее время в формате ANSIC
		req.Method,                    // метод запроса (GET, POST и т.д.)
		req.URL,                       // URL запроса
	)

	// Передаем запрос следующему обработчику
	return loggerRT.next.RoundTrip(req)
}
