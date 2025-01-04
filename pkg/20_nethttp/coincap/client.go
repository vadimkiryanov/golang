package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Client представляет клиент для работы с API CoinCap
type Client struct {
	client *http.Client
}

// Базовый URL для API CoinCap
var baseURL = "https://api.coincap.io/v2/assets"

// NewClient создает новый экземпляр клиента CoinCap
// timeout: максимальное время ожидания ответа от сервера
// возвращает ошибку, если timeout равен 0
func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("time duration is zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &LoggingRoundTripper{
				logger: io.Writer(os.Stdout),
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

// GetAssets получает список всех криптовалют
// isRedirect: флаг для переключения между HTTP и HTTPS
// возвращает массив данных о криптовалютах или ошибку
func (clientStruct *Client) GetAssets(isRedirect bool) ([]*AssetData, error) {
	// Переключение на HTTP если isRedirect true
	if isRedirect {
		baseURL = "http://api.coincap.io/v2/assets"
	}

	// Выполнение GET-запроса к API
	response, err := clientStruct.client.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close() // Гарантируем закрытие тела ответа

	fmt.Printf("response.Status: %v\n", response.Status)

	// Чтение тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err, body)
	}

	// Структура для десериализации JSON-ответа
	var assetsData *AssetsResponse

	// Преобразование JSON в структуру
	if err := json.Unmarshal(body, &assetsData); err != nil {
		log.Fatal(err)
	}

	// Вывод информации о каждой криптовалюте
	for _, elem := range assetsData.Assets {
		fmt.Println(elem.GetInfo())
	}

	return []*AssetData{}, nil
}

// GetAsset получает информацию о конкретной криптовалюте по её идентификатору
// idAsset: идентификатор криптовалюты
// возвращает данные о криптовалюте или ошибку
func (clientStruct *Client) GetAsset(idAsset string) (*AssetData, error) {
	// Формирование URL для конкретной криптовалюты
	var URL = fmt.Sprintf(baseURL+"/%s", idAsset)

	// Выполнение GET-запроса
	response, err := clientStruct.client.Get(URL)
	if err != nil {
		return &AssetData{}, err
	}
	defer response.Body.Close()

	fmt.Printf("response.Status: %v\n", response.Status)

	// Чтение тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err, body)
	}

	// Структура для десериализации JSON-ответа
	var assetData *AssetResponse

	// Преобразование JSON в структуру
	if err := json.Unmarshal(body, &assetData); err != nil {
		log.Fatal(err)
	}

	// Вывод информации о криптовалюте
	fmt.Println(assetData.Asset.GetInfo())

	return &AssetData{}, nil
}
