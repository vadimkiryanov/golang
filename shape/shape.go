package shape

import (
	"fmt"
	"math"
)

// ======= СТРУКТУРЫ =======
type Age int

func (a Age) IsAdult() bool {
	if a >= 18 {
		return true
	} else {
		return false
	}
}

// Структура
type User struct {
	Name   string
	Age    Age
	Gender string
	City   string
}

// Value resiver
// Расширение структуры, добавление метода "printUserInfo"
// копирует и не изменяет исходный объект
func (u User) GetName() string {
	return u.Name
}

// Pointer resiver
// Ссылается на исходный объект и меняет его
func (u *User) SetName(newName string) {
	u.Name = newName
}

// Структура с мапой
type DumbDatabase struct {
	mMap map[string]User
}

// ======= ИНТЕРФЕЙСЫ =======
type ShapeWithArea interface {
	GetArea() float32
}
type ShapeWithPerimeter interface {
	GetPerimeter() float32
}

// Общий интерфейс для фигур, которые будут иметь метод getArea() и getPerimeter()
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

// Структура Square (Квадрат) которая имеет поле sideLength
type Square struct {
	sideLength float32
}

func NewSquare(sideLength float32) Square {
	return Square{sideLength}
}

// Определение метода getArea для структуры Square
func (s Square) GetArea() float32 {
	return s.sideLength * s.sideLength
}

// Определение метода getPerimeter для структуры Square
func (s Square) GetPerimeter() float32 {
	return s.sideLength * 4
}

// Структура Circle (Круг)
type Circle struct {
	raduis float32
}

func NewCircle(raduis float32) Circle {
	return Circle{raduis}
}

// Определение метода getArea для структуры Circle
func (c Circle) GetArea() float32 {
	return c.raduis * c.raduis * math.Pi
}

// Определение метода getPerimeter для структуры Circle
func (c Circle) GetPerimeter() float32 {
	return c.raduis * 2 * math.Pi
}

// Функция принта метода getArea, сюда можно поместить любую структуру,
// которая имеет метод getArea
func PrintShapeArea(shape Shape) {
	var typeName string
	switch shape.(type) {
	case Square:
		typeName = "Square"
	case Circle:
		typeName = "Circle"
	}

	fmt.Println("Area: ", shape.GetArea(), typeName)
	fmt.Println("Perimeter: ", shape.GetPerimeter(), typeName)
}
