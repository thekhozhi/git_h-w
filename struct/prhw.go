package main

import "fmt"

// 1 exercise.

type Clients struct {
	Id     int
	Name   string
	Basket []CardProducts
}

type CardProducts struct {
	Name       string
	Product_id int
	Size_id    int
	Quantity   int
}

type Product struct {
	Id   int
	Name string
	size []Size
}

type Size struct {
	Id    int
	Name  string
	Price int
}

var clients = []Clients{
	{
		Id:   23,
		Name: "John",
		Basket: []CardProducts{
			{Product_id: 1, Size_id: 100, Quantity: 3},
			{Product_id: 2, Size_id: 200, Quantity: 2},
			{Product_id: 3, Size_id: 300, Quantity: 5},
		},
	},
	{
		Id:   17,
		Name: "Nick",
		Basket: []CardProducts{
			{Product_id: 1, Size_id: 100, Quantity: 15},
			{Product_id: 2, Size_id: 200, Quantity: 10},
			{Product_id: 3, Size_id: 300, Quantity: 8},
		},
	},
}

var product = []Product{
	{
		Id:   1,
		Name: "Pizza regular",
		size: []Size{
			{Id: 100, Name: "25sm", Price: 28000},
		},
	},
	{
		Id:   2,
		Name: "Pizza Peperoni",
		size: []Size{
			{Id: 200, Name: "30sm", Price: 37000},
		},
	},
	{
		Id:   3,
		Name: "Pizza Italian",
		size: []Size{
			{Id: 300, Name: "35sm", Price: 45000},
		},
	},
}

func SumOfProducts() {
	for _, vclients := range clients {
		sum := 0
		for _, vbasket := range vclients.Basket {
			for _, vproduct := range product {
				for _, vsize := range vproduct.size {
					if vbasket.Size_id == vsize.Id {
						sum = sum + (vsize.Price * vbasket.Quantity)
					}
				}
			}
		}
		fmt.Printf("%s'ni korzinasida %d sumlik product bor\n", vclients.Name, sum)
	}
}

func MaxProducts() {
	max := 0
	for _, vclients := range clients {
	    for _, vproduct := range product {
			for _, vsize := range vproduct.size {
				for _, vbasket := range vclients.Basket {
					if vbasket.Size_id == vsize.Id {
						if vbasket.Quantity > max {
							max = vbasket.Quantity
						}
					}
				}
			}
		}
		fmt.Println(vproduct.Name, max, "'ta")
	}
}

func main() {
	SumOfProducts()
	MaxProducts()
}
