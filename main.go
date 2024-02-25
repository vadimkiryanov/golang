package main

import (
	"fmt"
	"reflect"
)

func main() {

	var messageInit string
	messageInit = "World!!!"

	fmt.Println(reflect.TypeOf(messageInit), messageInit)
}
