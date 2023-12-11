package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Branch struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Transaction struct {
	Id        int    `json:"id"`
	BranchId  int    `json:"branch_id"`
	ProductId int    `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
}

var branches = []Branch{}
var transactions = []Transaction{}
var categories = []Category{}
var products = []Product{}

func WorkWithJson() bool {
	// 1
	branchfile, err := os.OpenFile("branches.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error while opening branches.json file!", err)
	}

	err = json.NewDecoder(branchfile).Decode(&branches)
	if err != nil {
		fmt.Println("Error while decoding branches", err)
	}

	// 2
	transfile, err := os.OpenFile("branch_transaction.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error while opening transaction file!", err)
	}

	err = json.NewDecoder(transfile).Decode(&transactions)
	if err != nil {
		fmt.Println("Error while decoding transactions", err)
	}

	// 3
	categoryfile, err := os.OpenFile("categories.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error while opening category file!", err)
	}

	err = json.NewDecoder(categoryfile).Decode(&categories)
	if err != nil {
		fmt.Println("Error while decoding categories", err)
	}

	// 4
	productfile, err := os.OpenFile("products.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error while opening product json file!", err)
	}

	err = json.NewDecoder(productfile).Decode(&products)

	return true
}

func main() {
	WorkWithJson()

	// 1 TASK
	// 	for _, vcategory := range categories{
	// 		counter := 0
	// 		for _, vtrans := range transactions{
	// 			if vtrans.ProductId == vcategory.Id{
	// 				counter++
	// 			}
	// 		}
	// 		fmt.Println(vcategory.Name,"->",counter)
	// 	}

	// categoryMap := make(map[int]string)
	// productMap := make(map[int]int)
	// branchMap := make(map[int]string)

	// for _, vbr := range branches{
	// 	branchMap[vbr.Id] = vbr.Name
	// }
	//fmt.Println(branchMap)

	// for _, vcat := range categories{
	// 	categoryMap[vcat.Id] = vcat.Name
	// }
	//fmt.Println(categoryMap)

	// for _, v := range categories {
	// 	productMap = categoryMap[v.Id]++
	// }
	// fmt.Println(productMap)

	// 2 TASK

	// for _, vbranch := range branches{
	// 	counter := 0
	// 	for _, vcat := range categories{
	// 		for _, vtrans := range transactions{
	// 			if vbranch.Id == vtrans.BranchId{
	// 				if vtrans.ProductId == vcat.Id{
	// 					counter++
	// 				}
	// 			}
	// 		}
	// 		fmt.Println(vbranch.Name,"\n",vcat.Name,"->",counter)
	// 	}
	// }

	branchMap := make(map[int]string)
	counterMap := make(map[int]int)

	for _, v := range branches {
		branchMap[v.Id] = v.Name
	}

	for _, v := range transactions {
		counterMap[v.BranchId] = counterMap[v.BranchId] + 1
	}
	fmt.Println(counterMap)
	
	type Count struct{
		Key int
		Value int
	}
	counts := []Count{}

	for i, v := range counterMap{
		counts = append(counts, Count{i,v})
	}
	fmt.Println(counts)

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})

	fmt.Println(counts)
}
