package main

import (
	"fmt"
)

// task 1
type Animal struct{ Name string }
type Dog struct {
	Animal
	Pron string
}
type Cat struct {
	Animal
	Pron string
}
type Saying interface{ Say() }

func (cat Cat) Say() {
	fmt.Printf("%v, меня зовут %v!\n", cat.Pron, cat.Name)
}
func (dog Dog) Say() {
	fmt.Printf("%v, меня зовут %v!\n", dog.Pron, dog.Name)
}

// task2
type Engine struct{ Power uint }
type Car struct{ Engine }
type Motorcycle struct{ Engine }

// task3
type User struct{ Login, Password string }
type Profile struct {
	User
	Email string
}

// task4
type Shape struct{}

func (shape Shape) Draw() {}

type Circle struct{ Shape }
type Rectangle struct{ Shape }

// @Override
func (circle Circle) Draw() { 
	fmt.Println("чиркаем карандашем и кружочек")
}
func (rectangle Rectangle) Draw() {
	fmt.Println("две попарно параллельных прямых, и вот это чудо, прямоугольник")
}

// task5
type Weapon struct{}

func (weapon Weapon) Shoot() {}

type Pistol struct{ Weapon }

func (pistol Pistol) Shoot() {
	fmt.Printf("Курок спускается с пружины по нажатию спускового крючка, и боёк попадает по капсулю патрона калибра 9x19mm, который...\nв свою очередь из за удара взрывается, и воспламеняет порох внутри патрона, и газы...\nвыталкивают пулю с невероятной скоростью, и остальные газы приводят в действие blow-back систему...\nиз за которой гильза выкидывается из заслонки, и в патронник засылается новый патрон из магазина.\nУбойная сила отдачи распространилась по корпусу пистолета, что приведет к покачиванию кисти\n")
}

type Rifle struct{ Weapon }

func (rifle Rifle) Shoot() {
	fmt.Printf("Курок спускается с пружины по нажатию спускового крючка, и боёк попадает по капсулю патрона калибра 7.62х51mm, который...\nв свою очередь из за удара взрывается, и воспламеняет порох внутри патрона, и газы...\nвыталкивают патрон с невероятной скоростью, и остальные газы приводят в действие blow-back систему...\nиз за которой гильза выкидывается из заслонки, и в патронник засылается новый патрон из рожкового магазина АКС-74У.\nОставшиеся газы уйдут в газовую камеру, и в еще один канал с трубкой для газов (выйдет из пламегасителя)\n")
}

// task6
type Employee struct{}

func (employee Employee) Work() {}

type Manager struct{ Employee }
type Developer struct{ Employee }

func (dev Developer) Work() {
	fmt.Println("ыыыыы if err != nil {log.Fatal(err)}")
}
func (mgr Manager) Work() {
	fmt.Println("здравствуйте я менеджер анатолий")
}

// task7
type Switchable struct{}
type Adjustable struct{} 
type SmartLamp struct {
	Switchable
	Adjustable
}

// task10
type Book struct {
	Title, Author string
}
type LibraryBook struct {
	ISBN      string
	Available bool
}

func main() {
	// task1на этот раз сразу с реализацией
	kotik := Cat{Animal{"Барсик"}, "Мяу-мяу"}
	sharik := Dog{Animal{"Шарик"}, "Гав-гав"}
	kotik.Say()
	sharik.Say() // ура работает
	//task2
	daihatsu_cuore := Car{Engine{45}}
	kawasaki_zzr1400 := Motorcycle{Engine{87}}
	_ = daihatsu_cuore
	_ = kawasaki_zzr1400
	//task3
	va1k0inen := Profile{User{"va1k0inen", "superpassword"}, "altushki@love4.ever"}
	_ = va1k0inen
	//task4
	newcircle := Circle{}
	newrectangle := Rectangle{}
	newcircle.Draw()
	newrectangle.Draw()
	//task5
	aks74u := Rifle{}
	glock26 := Pistol{}
	aks74u.Shoot()
	glock26.Shoot()
	//task6
	gopher := Developer{}
	taskmgr := Manager{}
	gopher.Work()
	taskmgr.Work()
	//task7
	smartlamp := SmartLamp{}
	_ = smartlamp
	// task14
	/*По факту, должен вызваться метод той встроенной структуры, которая была ближе всего при поиске во время обьявления элемента структуры.
	Здесь подойдет аналогия из DFS, где всегда идет учет на то, чтобы всегда найти кто ближе всего.*/
