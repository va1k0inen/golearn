package main

import (
	"fmt"
	"strings"
)

// task 1
type Point struct {
	X, Y int
}

func (uas Point) Display() {
	fmt.Printf("Точка: (%v, %v)", uas.X, uas.Y)
}

// task 2
type Student struct {
	Name string
	Age  int
}

func (tadzhiki Student) Introduce() {
	fmt.Printf("Меня зовут %v, мне %v лет", tadzhiki.Name, tadzhiki.Age)
}

// task 3
type Calculator struct {
	Result int
}

func (calc Calculator) Add(x int) {
	calc.Result += x
}

// task 4
type Counter struct {
	Value int
}

func (h *Counter) Increment() {
	h.Value++
}
func (a *Counter) Decrement() {
	a.Value--
}
func (asf *Counter) Reset() {
	asf.Value = 0
}

// task 5
type User struct {
	Username, Email string
}

func (ag *User) ChangeUsername(newName string) {
	ag.Username = newName
}
func (gaga *User) ChangeEmail(newEmail string) {
	gaga.Email = newEmail
}
func (ajv User) Display() {
	fmt.Printf("Адрес электронной почты: %v\nЮзернейм: %v\n", ajv.Email, ajv.Username)
}

// taask 6
type Cart struct {
	Items []string
}

func (sharik *Cart) AddItem(item string) {
	sharik.Items = append(sharik.Items, item)
}
func (uzbekcha *Cart) RemoveItem() {
	if len(uzbekcha.Items) > 0 {
		uzbekcha.Items = uzbekcha.Items[:len(uzbekcha.Items)-1]
	} else {
		fmt.Println("Корзина пуста")
	}
}
func (shat Cart) ShowItems() {
	for i, value := range shat.Items {
		fmt.Printf("Товар №%v: %v", i, value)
	}
}

// task 7
type Validator struct{}

func (valid Validator) IsEmailValid(email string) bool {
	if strings.Contains(email, "@") {
		return true
	} else {
		return false
	}
}

// task 8
type Temperature struct {
	Celcius float64
}

func (tempx Temperature) ToFahrenheit() float64 {
	return tempx.Celcius*1.8 + 32
}
func (tempy Temperature) ToKelvin() float64 {
	return tempy.Celcius + 273.5
}

// task 9
type StringProcessor struct{}

func (spx StringProcessor) UpperCase(s string) string {
	return strings.ToUpper(s)
}
func (spy StringProcessor) LowerCase(s string) string {
	return strings.ToLower(s)
}

// final task 10
type Test struct {
	Value int
}

func (testx *Test) ChangeByPointer(newValue int) {
	testx.Value = newValue
}
func (testy Test) ChangeByValue(newValue int) {
	testy.Value = newValue
}

func main() {
	xtest := Test{}
	xtest.ChangeByValue(10)

	ytest := Test{}
	ytest.ChangeByPointer(10)
	fmt.Printf("Меняем по значениююю: %v\tТеперь по указателю: %v", xtest.Value, ytest.Value)
}
