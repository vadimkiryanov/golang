package goroutine

import (
	"fmt"
	"time"
)

func SleepRun() {
	go fmt.Println("first")
	go fmt.Println("second")
	go fmt.Println("three")

	time.Sleep(time.Second)

	fmt.Println("FINISHED")
}
