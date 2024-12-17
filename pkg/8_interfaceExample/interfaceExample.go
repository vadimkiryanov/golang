package interfaceExample

import (
	"fmt"
	"main-mode/pkg/shape"
)

func Example() {
	var square = shape.NewSquare(50)
	var circle = shape.NewCircle(50)

	// Пустой интерфейс
	var emptyInterface = func(i interface{}) {
		// Проверка типа
		switch value := i.(type) {
		case int:
			fmt.Println("int", value)
		case string:
			fmt.Println("string", value)
		default:
			fmt.Println("unknown type", value)
		}

		// Проверка на существование типа (type guard)
		var str, ok = i.(string)
		if !ok {
			fmt.Println("interface is not a string")
			return
		}

		fmt.Println("interface: ", len(str))
	}

	// logs
	shape.PrintShapeArea(square)
	shape.PrintShapeArea(circle)

	emptyInterface(2222)
	emptyInterface("circle")

}
