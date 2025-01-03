package nethttp

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (loggerRT loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(loggerRT.logger, "=== loggingRoundTripper: [%s] %s %s \n", time.Now().Format(time.ANSIC), req.Method, req.URL)
	return loggerRT.next.RoundTrip(req)
}

func Main() {
	// Инициализация клиента
	var client = &http.Client{
		// Хэндлер для перенаправления
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Printf("\"REDIRECT\": %v\n", "REDIRECT")
			return nil
		},

		// Мини мидлвэйр
		Transport: &loggingRoundTripper{
			logger: io.Writer(os.Stdout),
			next:   http.DefaultTransport,
		},

		// Кастомное ожидание ответа сервера
		Timeout: 5 * time.Second,
	}

	// Запрос
	var response, err = client.Get("http://example.com/")

	// Обработка ошибки
	if err != nil {
		log.Fatal(err)
	}

	// Закрытие ответа после использования
	defer response.Body.Close()

	fmt.Printf("response.Status: %v\n", response.Status)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err, body)
	}

	fmt.Printf("body: %v\n", string(body))
}
