package main

// --PRACTICE--

// 1 exercise.

// package main

// import "fmt"

// type Car struct {
// 	Name       string
// 	Price      float64
// 	Model      string
// 	HorsePower int
// }

// func main() {
// 	cars := []Car{

// 		Car{

// 			Name:       "sdfw",
// 			Price:      15000,
// 			Model:      "x5",
// 			HorsePower: 280,
// 		},
// 		Car{

// 			Name:       "bmw",
// 			Price:      15000,
// 			Model:      "m5",
// 			HorsePower: 280,
// 		},
// 		Car{
// 			Name:       "merc",
// 			Price:      17000,
// 			Model:      "amg",
// 			HorsePower: 320,
// 		},
// 	}
// 	for i := 0; i < len(cars); i++ {
// 		for j := 0; j < len(cars); j++ {
// 			if cars[i].Model < cars[j].Model {
// 				cars[i].Model, cars[j].Model = cars[j].Model, cars[i].Model
// 			}
// 		}
// 	}

// 	for _, v := range cars {
// 		fmt.Println(v)
// 	}
// }

//2 exercise.

// package main

// import "fmt"

// type Team struct{

// 	Name string
// 	Coach string
// 	PlayersCount uint
// 	Captain string
// }

// func main(){

// 	team := []Team{
// 		{
// 			Name: "Real Madrid",
// 			Coach: "Ancelotti",
// 			PlayersCount: 15,
// 			Captain: "Sergio Ramos",
// 		},
// 		{
// 			Name: "Barcelona",
// 			Coach: "Havi",
// 			PlayersCount: 20,
// 			Captain: "Ter Stegen",
// 		},
// 		{
// 			Name: "Liverpool",
// 			Coach: "Jurgen Clopp",
// 			PlayersCount: 17,
// 			Captain: "Virjil Vandijk",
// 		},
// 	}

// for i := 0; i < len(team); i++{
// 	for j := 0; j<len(team); j++{
// 		 if team[i].PlayersCount > team[j].PlayersCount{
// 			team[i],team[j] = team[j],team[i]
// 		 }
// 	}
// }

// for _, value := range(team){
// 	fmt.Println(value)
// }
// }

// 3 exercise.

// package main

// import "fmt"

// type Student struct {
// 	Name        string
// 	Age         int
// 	Scholarship int
// 	Score       []int
// }

// func main() {
// 	s := []Student{
// 		Student{
// 			Name:        "Tom",
// 			Age:         20,
// 			Scholarship: 10,
// 			Score:       []int{4, 5, 4, 5, 4},
// 		},
// 		Student{
// 			Name:        "Tommy",
// 			Age:         20,
// 			Scholarship: 10,
// 			Score:       []int{5, 5, 5, 5, 5},
// 		},
// 		Student{
// 			Name:        "Hang",
// 			Age:         20,
// 			Scholarship: 10,
// 			Score:       []int{4, 5, 3, 3, 4},
// 		},
// 	}

// 	for i := 0; i < len(s); i++{
// 		student := s[i]
// 		sum := 0
// 		for j := 0; j < len(student.Score); j++{
// 			sum = sum + student.Score[j]
// 		}

// 		if sum / len(student.Score) == 4{
// 			fmt.Println(student)
// 		}
// 	}
// }

//--HOMEWORK--.

// 1 exercise.

// package main

// import (
// 	"fmt"
// )

// type Student struct{
// 	Name string
// 	Scholarship uint
// 	Course uint
// }

// func main(){
// s := []Student{
// 	{
// 		Name: "Tom",
// 		Scholarship: 150000,
// 		Course: 3,
// 	},
// 	{
// 		Name: "Candy",
// 		Scholarship: 95000,
// 		Course: 1,
// 	},
// 	{
// 		Name: "Tommy",
// 		Scholarship: 270000,
// 		Course: 4,
// 	},
// 	{
// 		Name: "Jack",
// 		Scholarship: 120000,
// 		Course: 2,
// 	},
// 	{
// 		Name: "Sam",
// 		Scholarship: 200000,
// 		Course: 4,
// 	},
// 	{
// 		Name: "David",
// 		Scholarship: 125000,
// 		Course: 2,
// 	},
// 	{
// 		Name: "Amelia",
// 		Scholarship: 147000,
// 		Course: 3,
// 	},
// 	{
// 		Name: "Travers",
// 		Scholarship: 97000,
// 		Course: 1,
// 	},
// 	{
// 		Name: "Kate",
// 		Scholarship: 220000,
// 		Course: 4,
// 	},
// 	{
// 		Name: "John",
// 		Scholarship: 118000,
// 		Course: 2,
// 	},
// }
// sum := 0;
// for i := 0; i < len(s);i++{
// 	if s[i].Course == 2{
// 		sum = sum + int(s[i].Scholarship)
// 	}
// }
// fmt.Printf("Umumiy hisobda 2 kurs studentlariga %d tolandi",sum)
// fmt.Println()
// }

// 2 exercise.

// package main

// import (
// 	"fmt"
// )

// type Student struct{
// 	Name string
// 	Scholarship uint
// 	Course uint
// }

// func main(){
// s := []Student{
// 	{
// 		Name: "Tom",
// 		Scholarship: 150000,
// 		Course: 3,
// 	},
// 	{
// 		Name: "Candy",
// 		Scholarship: 95000,
// 		Course: 1,
// 	},
// 	{
// 		Name: "Tommy",
// 		Scholarship: 270000,
// 		Course: 4,
// 	},
// 	{
// 		Name: "Jack",
// 		Scholarship: 120000,
// 		Course: 2,
// 	},
// 	{
// 		Name: "Sam",
// 		Scholarship: 200000,
// 		Course: 4,
// 	},
// 	{
// 		Name: "David",
// 		Scholarship: 125000,
// 		Course: 2,
// 	},
// 	{
// 		Name: "Amelia",
// 		Scholarship: 147000,
// 		Course: 3,
// 	},
// 	{
// 		Name: "Travers",
// 		Scholarship: 97000,
// 		Course: 1,
// 	},
// 	{
// 		Name: "Kate",
// 		Scholarship: 220000,
// 		Course: 4,
// 	},
// 	{
// 		Name: "John",
// 		Scholarship: 118000,
// 		Course: 2,
// 	},
// }
// for i := 0; i < len(s); i++ {
// 	if len(s[i].Name) < 5 {
// 		fmt.Println(s[i].Name)
// 	}
// }
// }

// 3 exercise.

// package main

// import "fmt"

// type Aeroport struct {

// 	Planetype string
// 	FlightDate string
// 	FromCity string
// 	ToCity string
// 	FlightTime float32
// }

// func main(){

// 	aero := []Aeroport{
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Spring",
// 			FromCity: "New York",
// 			ToCity: "Istanbul",
// 			FlightTime: 3.15,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Autumn",
// 			FromCity: "New York",
// 			ToCity: "Tashkent",
// 			FlightTime: 2.30,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Summer",
// 			FromCity: "New York",
// 			ToCity: "Mekka",
// 			FlightTime: 9.25,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Winter",
// 			FromCity: "New York",
// 			ToCity: "Moscow",
// 			FlightTime: 11.39,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Autumn",
// 			FromCity: "New York",
// 			ToCity: "Tashkent",
// 			FlightTime: 2.15,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Spring",
// 			FromCity: "New York",
// 			ToCity: "Istanbul",
// 			FlightTime: 5.40,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Summer",
// 			FromCity: "New York",
// 			ToCity: "Mekka",
// 			FlightTime: 6.30,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Autumn",
// 			FromCity: "New York",
// 			ToCity: "Moscow",
// 			FlightTime: 1.20,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Summer",
// 			FromCity: "New York",
// 			ToCity: "Tashkent",
// 			FlightTime: 2.40,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Spring",
// 			FromCity: "New York",
// 			ToCity: "Istanbul",
// 			FlightTime: 3.58,
// 		},
// 	}
// 	for i := 0; i < len(aero); i++ {
// 		if aero[i].FlightDate == "Summer"{
// 			fmt.Println(aero[i])
// 		}
// 	}
// }

// 4 exercise.

// package main

// import "fmt"

// type Aeroport struct {

// 	Planetype string
// 	FlightDate string
// 	FromCity string
// 	ToCity string
// 	FlightTime float32
// }

// func main(){

// 	aero := []Aeroport{
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Spring",
// 			FromCity: "New York",
// 			ToCity: "Istanbul",
// 			FlightTime: 3.15,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Autumn",
// 			FromCity: "New York",
// 			ToCity: "Tashkent",
// 			FlightTime: 4.35,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Summer",
// 			FromCity: "New York",
// 			ToCity: "Mekka",
// 			FlightTime: 9.25,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Winter",
// 			FromCity: "New York",
// 			ToCity: "Moscow",
// 			FlightTime: 11.39,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Autumn",
// 			FromCity: "New York",
// 			ToCity: "Tashkent",
// 			FlightTime: 2.15,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Spring",
// 			FromCity: "New York",
// 			ToCity: "Istanbul",
// 			FlightTime: 5.40,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Summer",
// 			FromCity: "New York",
// 			ToCity: "Mekka",
// 			FlightTime: 6.30,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Autumn",
// 			FromCity: "New York",
// 			ToCity: "Moscow",
// 			FlightTime: 1.20,
// 		},
// 		{
// 			Planetype: "Boing 737",
// 			FlightDate: "Summer",
// 			FromCity: "New York",
// 			ToCity: "Tashkent",
// 			FlightTime: 2.40,
// 		},
// 		{
// 			Planetype: "Boing 780",
// 			FlightDate: "Spring",
// 			FromCity: "New York",
// 			ToCity: "Istanbul",
// 			FlightTime: 3.58,
// 		},
// 	}
// 	for i := 0; i < len(aero); i++ {
// 		if aero[i].FlightTime > 2.00 && aero[i].FlightTime < 3.00 && aero[i].ToCity == "Tashkent"{
// 			fmt.Println(aero[i])
// 		}
// 	}
// }

// 5 exercise.

// package main

// import "fmt"

// type Team struct {
// 	Name         string
// 	Coach        string
// 	PlayersCount uint
// 	Captain      string
// }

// func main() {

// 	team := []Team{
// 		{
// 			Name:         "RealMadrid",
// 			Coach:        "Ancelotti",
// 			PlayersCount: 15,
// 			Captain:      "Sergio Ramos",
// 		},
// 		{
// 			Name:         "Barcelona",
// 			Coach:        "Havi",
// 			PlayersCount: 20,
// 			Captain:      "Ter Stegen",
// 		},
// 		{
// 			Name:         "Liverpool",
// 			Coach:        "Jurgen Clopp",
// 			PlayersCount: 25,
// 			Captain:      "Virjil Vandijk",
// 		},
// 	}
// 	var name string
// 	fmt.Print("Nameni kiriting: ")
// 	fmt.Scan(&name)

// 	for i := 0; i < len(team); i++ {
// 		if name == team[i].Name {
// 			fmt.Println(team[i])
// 		}
// 	}
// 	for i := 0; i < len(team); i++ {
// 		for j := 0; j < len(team); j++ {
// 			if team[i].PlayersCount > team[j].PlayersCount {
// 				team[i], team[j] = team[j], team[i]
// 			}
// 		}
// 	}
// 	fmt.Println(team)
// }
