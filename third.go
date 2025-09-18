package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

// task1
type Speaker interface{ Speak() string }
type Dog struct{}
type Cat struct{}
type Cow struct{}

func (d Dog) Speak() string   { return "гав гав" }
func (c Cat) Speak() string   { return "мяяуу" }
func (cow Cow) Speak() string { return "муууу" }

// task2
type Calculator interface{ Calculate(a, b int) int }
type Adder struct{}
type Subtractor struct{}
type Multiplyer struct{}

func (add Adder) Calculate(a, b int) int       { return a + b }
func (sub Subtractor) Calculate(a, b int) int  { return a - b }
func (mltp Multiplyer) Calculate(a, b int) int { return a * b }

// task3
type Vehicle interface{ Move() string }
type Car struct{}
type Bike struct{}
type Plane struct{}

func (car Car) Move() string {
	return "11000 RPM 3 GEAR 140 KM/h"
}
func (bike Bike) Move() string { return "13000 RPM 6 GEAR 310 KM/H" }
func (plane Plane) Move() string {
	return "-я стал пилотом, чтобы побороть свой страх / -серьезно? а чего ты боишься? летать в самолете? / -страх смерти в одиночестве"
}

// task4
type Logger interface{ Log(message string) }
type ConsoleLogger struct{}
type FileLogger struct{}

func (cl ConsoleLogger) Log(message string) { log.Println(message) }
func (fl FileLogger) Log(message string) {
	data := []byte(message)
	err := os.WriteFile("cmFuZG9tbHkg.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// task5
type Shape interface{ Area() float64 }
type Circle struct{ Radius float64 }
type Rectangle struct{ X, Y float64 }

func (circle Circle) Area() float64       { return math.Pi * circle.Radius * circle.Radius }
func (rectangle Rectangle) Area() float64 { return rectangle.X * rectangle.Y }

// task6
type Sorter interface{ Sort(arr []int) []int }
type BubbleSorter struct{}
type QuickSorter struct{}

func (bs BubbleSorter) Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
func (qs QuickSorter) Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	var left, right []int
	for _, x := range arr[1:] {
		if x < pivot {
			left = append(left, x)
		} else {
			right = append(right, x)
		}
	}

	return append(append(qs.Sort(left), pivot), qs.Sort(right)...)
}

// task7
type Switchable interface {
	TurnOn()
	TurnOff()
}
type Adjustable interface {
	SetLevel(perc int) int
}
type Light struct {
	IsTurn  bool
	Percent uint8
}

func (light *Light) TurnOn()               { light.IsTurn = true }
func (light *Light) TurnOff()              { light.IsTurn = false }
func (light *Light) SetLevel(perc int) int { light.Percent = uint8(perc); return int(light.Percent) }

type Thermostat struct {
	IsTurn bool
	Temp   int
}

func (thermo *Thermostat) TurnOn()               { thermo.IsTurn = true }
func (thermo *Thermostat) TurnOff()              { thermo.IsTurn = false }
func (thermo *Thermostat) SetLevel(perc int) int { thermo.Temp = perc; return thermo.Temp }

// task8
type Mover interface{ Move() }
type Attacker interface{ Attack() }
type Talker interface{ Talk() }
type Hero struct{}
type Monster struct{}

func (hero Hero) Move() {
	fmt.Println("*звук меча, ерзающий по кольчуге и по броне")
} // хардкодим ребята!!
func (monster Monster) Move() {
	fmt.Println("*шорох листьев, громкий звук топота")
}
func (hero Hero) Talk()       { fmt.Println("painu helvettiin vittun perkele saatana") }
func (monster Monster) Talk() { fmt.Println("*неразборчиво") }
func (hero Hero) Attack() {
	fmt.Println("достает меч и наносит удар острием *MISS")
}
func (monster Monster) Attack() {
	fmt.Println("замахивается рукой и размазывает персонажа")
}

func main() {

}
