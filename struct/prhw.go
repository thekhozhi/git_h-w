package main

// 1 exercise.
//
// package main

// type Clients struct {
// 	Id     int
// 	Name   string
// 	Basket []CardProducts
// }

// type CardProducts struct {
// 	Product_id int
// 	Size_id    int
// 	Quantity   int
// }

// type Product struct {
// 	Id   int
// 	Name string
// 	size []Size
// }

// type Size struct {
// 	Id    int
// 	Name  string
// 	Price int
// }

// var clients = []Clients{
// 	{
// 		Id:   23,
// 		Name: "John",
// 		Basket: []CardProducts{
// 			{Product_id: 1, Size_id: 100, Quantity: 3},
// 			{Product_id: 2, Size_id: 200, Quantity: 2},
// 			{Product_id: 3, Size_id: 300, Quantity: 5},
// 		},
// 	},
// 	{
// 		Id:   17,
// 		Name: "Nick",
// 		Basket: []CardProducts{
// 			{Product_id: 1, Size_id: 100, Quantity: 15},
// 			{Product_id: 2, Size_id: 200, Quantity: 10},
// 			{Product_id: 3, Size_id: 300, Quantity: 8},
// 		},
// 	},
// }

// var product = []Product{
// 	{
// 		Id:   1,
// 		Name: "Pizza regular",
// 		size: []Size{
// 			{Id: 100, Name: "25sm", Price: 28000},
// 		},
// 	},
// 	{
// 		Id:   2,
// 		Name: "Pizza Peperoni",
// 		size: []Size{
// 			{Id: 200, Name: "30sm", Price: 37000},
// 		},
// 	},
// 	{
// 		Id:   3,
// 		Name: "Pizza Italian",
// 		size: []Size{
// 			{Id: 300, Name: "35sm", Price: 45000},
// 		},
// 	},
// }

// func main() {
// 	for _, vclients := range clients {
// 		sum := 0
// 		for _, vCardProducts := range vclients.Basket {
// 			for _, vProduct := range product {
// 				if vCardProducts.Product_id == vProduct.Id {
// 						for _, vSize := range vProduct.size {
// 						sum = sum + (vSize.Price*vCardProducts.Quantity)
// 					}
// 				}
// 			}
// 		}
// 	}
// }
