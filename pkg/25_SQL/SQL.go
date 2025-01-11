package sql_pkg

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func SQL() {
	// Подключение к БД
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable password=qwerty123")

	// Проверка подключения
	if err != nil {
		log.Fatal(err)
	}

	// Закрытие соединения
	defer db.Close()

	// Проверка соединения
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Вывод сообщения о подключении
	fmt.Println("Connected to database")

	// Добавление пользователя
	if err := insertUser(db, User{
		Name:     "asdqqqqSSSSJohn asdasdDoe",
		Email:    "j@j.com",
		Password: "123",
	}); err != nil {
		log.Fatal(err)
	}

	// Удаление пользователя
	if err := deleteUser(db, 3); err != nil {
		log.Fatal(err)
	}

	// Обновление пользователя
	if err := updateUser(db, 5, User{
		Name:  "Doe Doe222",
		Email: "j@j123321.com",
	}); err != nil {
		log.Fatal(err)
	}

	users, err := getUsers(db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No rows")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(users)

	user, err := getUserById(db, 1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No rows")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(user)

}

// SQL запрос на получение всех пользователей
func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0)

	// Получение данных
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisteredAT)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("Selected all users...")
	return users, nil
}

// SQL запрос на получение одного пользователя
func getUserById(db *sql.DB, id int64) (User, error) {
	var user User
	err := db.QueryRow("select * from users where id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisteredAT)

	fmt.Println("Selected user by id...")
	return user, err
}

func insertUser(db *sql.DB, user User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	_, err = tx.Exec("insert into users (name, email, password) values ($1, $2, $3)",
		user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into logs (entity, action) values ($1, $2)",
		"user", "created")
	if err != nil {
		return err
	}

	fmt.Println("Inserted...")

	return tx.Commit()
}

func deleteUser(db *sql.DB, id int64) error {
	_, err := db.Exec("delete from users where id = $1", id)

	fmt.Println("Deleted...")
	return err
}

func updateUser(db *sql.DB, id int, updUser User) error {
	_, err := db.Exec("update users set name = $1, email = $2 where id = $3",
		updUser.Name, updUser.Email, id)

	fmt.Println("Updated...")
	return err
}

type User struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	RegisteredAT time.Time
}
