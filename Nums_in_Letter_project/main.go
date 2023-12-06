package main

import "fmt"

func NumsInLetter(num int) {
	n := 0
	tempNum2 := num
	tempNum1 := num
	temp := 0
	ten := 1
	for tempNum1 != 0 {
		tempNum1 = tempNum1 / 10
		temp++
	}
	for i := 1; i < temp; i++ {
		ten *= 10
	}
	n = tempNum2 / ten
	switch n {
		case 1:
			fmt.Print("bir ")
		case 2:
			fmt.Print("ikki ")
		case 3:
			fmt.Print("uch ")
		case 4:
			fmt.Print("tort ")
		case 5:
			fmt.Print("besh ")
		case 6:
			fmt.Print("olti ")
		case 7:
			fmt.Print("yetti ")
		case 8:
			fmt.Print("sakkiz ")
		case 9: 
			fmt.Print("toqqiz ")
	}

	
}

func main() {
	NumsInLetter(123456)
}
