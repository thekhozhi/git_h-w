package main

// 1 exercise.

// package main

// import (
//     "fmt"
// 	 "os"
// )

// func main(){
// var m = map[string]string{
// 	"english": "iglizcha",
// 	"name": "ism",
// 	"age": "yosh",
// 	"pen": "ruchka",
// 	"board": "doska",
// 	"bag": "sumka",
// 	"dad": "ota",
// 	"desk": "parta",
// 	"fear": "qorqinch",
// 	"victory": "galaba",

// }
// file, err := os.OpenFile("file.txt", os.O_RDWR | os.O_CREATE, 0666)
// if err != nil{
// 	panic(err)
// }
// defer file.Close()

// var count int
// for i,v := range m {
// 	count++
//  _, err := file.WriteString(fmt.Sprintf("%d. %s - %s\n", count, i, v))
// 	 if err != nil {
// 		panic(err)
// 	 }
// 	}

// }

// 2 exercise.

// type Company struct {
// 	Id      int       `json:"id"`
// 	Name    string    `json:"name"`
// 	Owner   Owner     `json:"owner"`
// 	Workers []Workers `json:"workers"`
// }

// type Workers struct {
// 	Id        int    `json:"id"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Salary    int    `json:"salary"`
// 	Level     string `json:"level"`
// }

// type Owner struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// }

/* func main() {
	file, err := os.Open("file.json")
	if err != nil {
		fmt.Println("Error while opening json file")
	}
	company := []Company{}
	err = json.NewDecoder(file).Decode(&company)
	if err != nil {
		fmt.Println("Error while decoding", err)
	}

	// 1 task
	// sort.Slice(company, func(i int, j int) bool{
	// 	return len(company[i].Workers) > len(company[j].Workers)
	// })
	// 	fmt.Println(company)

	// 2 task
	// 	wslice := []Workers{}
	// 	for _,comp := range company{
	// 		for _, worker := range comp.Workers{
	// 			wslice = append(wslice, worker)
	// 		}
	// 	}main
	// 	sort.Slice(wslice, func(i, j int) bool {
	// 		return wslice[i].Salary > wslice[j].Salary
	// 	})

	// 	for i, worker := range wslice {
	// 		fmt.Println(worker)
	// 		if i == 2{
	// 			break
	// 		}
} */

//
