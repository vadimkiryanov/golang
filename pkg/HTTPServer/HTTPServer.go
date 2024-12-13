package HTTPServer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

