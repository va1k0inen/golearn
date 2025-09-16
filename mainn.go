package main

// задача 1
type Point struct {
	X, Y int
}

// задача 2
type User struct {
	Name   string
	Age    int
	Active bool
}

// задача 3
type Book struct {
	Title, Author string
	Year          int
}

func main() {
	// задача 4
	A := Point{0, 0}
	B := Point{5, 3}
	C := Point{-2, 7}
	// задача 5
	serega := User{"Сергей", 13, true}
	andryha := User{"Андрей", 13, true}
	volodya := User{"Владимир", 14, true}
	// задача 6
	warjapeace := Book{"Война и мир", "Л. Толстой", 1869}
	nineteeneightyfour := Book{"1984", "Оруэлл", 1949}
	harrypotter := Book{"Гарри Поттер и философский камень", "Дж. Роулинг", 1997}
	// задача 7
	newpoint := Point{10, 20}
	newpoint.X = 15
	newpoint.Y = 25
	// задача 8
	newUser := User{"Михаил", 25, true}
	newUser.Age += 1
	// задача 9
	newBook := Book{"Старая книга", "null", 1465}
	newBook.Title = "Новая книга"
	// задача 10
	anondata := struct {
		Theme    string
		FontSize int
		DarkMode bool
	}{
		Theme:    "Matcha",
		FontSize: 12,
		DarkMode: true,
	}
	// задача 11
	anonweather := struct {
		Temperature int
		Weather     string
		Humidity    int
	}{
		Temperature: 18,
		Weather:     "Rainy",
		Humidity:    36,
	}
	// задача 12
	anonjson := struct {
		Status string
		Code   int
	}{
		Status: "OK",
		Code:   200,
	}

	_ = anonjson // чтобы компилятор не ругался
	_ = anonweather
	_ = anondata
	_ = A
	_ = B
	_ = C
	_ = serega
	_ = andryha
	_ = volodya
	_ = newBook
	_ = newUser
	_ = newpoint
	_ = warjapeace
	_ = harrypotter
	_ = nineteeneightyfour
}
