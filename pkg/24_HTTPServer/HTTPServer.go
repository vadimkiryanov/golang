package HTTPServer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Main() {
	fmt.Println("SERVER STARTED")

	http.HandleFunc("/users", authMiddlaware(loggerMiddlaware(handleUsers)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}
type ctxKeys struct {
	userID string
}

var (
	users         = []User{{1, "Tom"}, {2, "Bob"}}
	ctxKeysValues = ctxKeys{"user_id"}
)

func handleUsers(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		getUsers(w, req)
	case http.MethodPost:
		addUser(w, req)
	case http.MethodDelete:
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)

}
func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var getNewId = func() int {
		var lastId = users[len(users)-1].ID
		var newId = lastId + 1

		return newId
	}

	user.ID = getNewId()

	users = append(users, user)

	fmt.Printf("Добавлен новый user: %v\n", user)

}

func loggerMiddlaware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var idFromCtx = r.Context().Value(ctxKeysValues.userID) // Получение userID из контекста
		userID, ok := idFromCtx.(string)                        // Приведение типа к строке
		if !ok {
			fmt.Printf("[%s] %s - ошибка: userID не валидный\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("Request: [%s] '%s' от user: %s\n", r.Method, r.URL, userID)
		next.ServeHTTP(w, r)
	})
}
func authMiddlaware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var userID = req.Header.Get(ctxKeysValues.userID) // Получение userID из хидера запроса

		if userID == "" {
			fmt.Printf("[%s] %s - ошибка: userID не найден\n", req.Method, req.URL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var ctx = req.Context()                                    // Создание контекста
		ctx = context.WithValue(ctx, ctxKeysValues.userID, userID) // Добавление userID в контекст
		req = req.WithContext(ctx)                                 // Добавление контекста в запрос

		next.ServeHTTP(w, req)
	})
}
