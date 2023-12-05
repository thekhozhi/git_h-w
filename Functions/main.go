package main
// --FUNCTIONS--

// 1 exercise

// package main

// import "fmt"

// func tubSon(){
// 	var counter int
// 	var num int
// 	fmt.Print("enter the num: ")
// 	fmt.Scan(&num)

// 	for i:=num; i<num+1; i++ {
// 		for j := 2; j < i/2; j++{
// 			if i % j == 0{
// 				counter++
// 			}
// 		}
// 		if counter == 0 {
// 			fmt.Println("Tub")
// 		}else {
// 			fmt.Println("Tub emas!")
// 		}
// 	}
// }

// func main(){
// 	tubSon()
// }

// 2 exercise.

// package main

// import "fmt"

// func numReverse(N int){
// 	counter := 0
// 	tempNum := N
// 	tempNum2 := N
// 	tempNum3 := N
// 	for tempNum != 0{
// 		tempNum = tempNum / 10
// 		counter++
// 	}
// 	for i := 0; i < counter; i++ {
// 		tempNum3 = tempNum2 % 10
// 		tempNum2 = tempNum2 / 10
// 		fmt.Print(tempNum3, " ")
// 	}
// 	 fmt.Println()
// }

// func main(){
// numReverse(7002)
// }

// 3 exercise.

// package main

// import "fmt"

// func sumNum(N int){
// tempNum := N
// tempNum2 := N
// tempNum3 := N
// counter := 0
// sum := 0

// for tempNum != 0 {
// 	tempNum = tempNum / 10
// 	counter++
// }
// for i := 0; i < counter; i++ {
//    tempNum2 = tempNum3 % 10
//    tempNum3 = tempNum3 / 10
//    sum = sum + tempNum2
// }
// 	fmt.Println(sum)
// }

// func main(){
// sumNum(12345555)
// }

// 4 exercise.

// package main

// import "fmt"

// func sumNum2(N int){
// tempNum := N
// sum := 0
// for tempNum != 0 {
// 	sum = sum + tempNum
// 	tempNum = tempNum - 1
// }
// fmt.Println(sum)
// }

// func main(){
// sumNum2(7)
// }

// 5 exercise.

// package main

// import "fmt"

// func Increase(N1, N2 int){
// 	a := N1
// 	b := N2
// 	temp1 := N1
// 	temp2 := N2
// 	sum := 0
// 	if a < b {
// 		sum = a - 1
// 		for i := a; i <= temp2; i++{
// 			sum = sum + 1
// 			fmt.Print(sum," ")
// 		}
// 		fmt.Println()
// 	}else if a > b {
// 		sum = a+1
// 		for i := b; i <= temp1; i++{
// 			sum = sum - 1
// 			fmt.Print(sum, " ")
// 		}
// 		fmt.Println()
// 	}else {
// 		fmt.Println("a and b are equal!!!")
// 	}
// }

// func main(){
// Increase(15, 23)
// }

// 6 exercise.

// package main

// import "fmt"

// func juftSonlar(N int){
// temp := N
// for temp != 0{
// 	temp = temp - 1
// 	if temp > 0 && temp % 2 == 0{
// 		fmt.Print(temp, " ")
// 	}
// }
// fmt.Println()
// }
// func main(){
// juftSonlar(10)
// }

// 7 exercise.

// package main

// import "fmt"

// func middleNum(n1, n2, n3 int){
// a := n1
// b := n2
// c := n3

// if a > b && a < c || a > c && a < b {
// 	fmt.Println(a)
// }else if b > a && b < c || b < a && b > c {
// 	fmt.Println(b)
// }else {
// 	fmt.Println(c)
// }
// }

// func main(){
// middleNum(2,0,1)
// }

// 8 exercise.

// package main

// import "fmt"

// func sumTrue(N int){
//  tempNum := N
// tempNum2 := N
// tempNum3 := N
// counter := 0
// sum := 0

// for tempNum != 0 {
// 	tempNum = tempNum / 10
// 	counter++
// }
// for i := 0; i < counter; i++ {
//    tempNum2 = tempNum3 % 10
//    tempNum3 = tempNum3 / 10
//    sum = sum + tempNum2
// }
// 	if sum % 2 == 0{
// 		fmt.Println(true)
// 	}else {
// 		fmt.Println(false)
// 	}
// }

// func main(){
// sumTrue(45651)
// }

// 9 exercise.

// package main

// import "fmt"

// func divideTwonums(n1, n2 float32)float32{
// z := n1 / n2
// return z
// }

// func main(){
// fmt.Println(divideTwonums(2,3))
// }

// 10 exercise.

// package main

// import "fmt"

// func sumUntiln(N int)int{
// tempNum := N
// sum := 0
// for tempNum != 0 {
// sum = sum + tempNum
// tempNum = tempNum - 1
// }
// return sum
// }

// func main(){
// fmt.Println(sumUntiln(6))
// }

// 11 exercise.

// package main

// import "fmt"
// func orderNum(N int){
// num := N
// temp := N
// count := 0
// for temp != 0 {
// 	temp = temp / 10
// 	count++
// }
// var tempNum = 1
// for i := 0; i < count; i++{
// 	if i == 0{
// 		for j := 1; j < count; j++{
// 			tempNum = tempNum * 10
// 		}
// 	}
// 	fmt.Print(num/tempNum," ")
// 	num %= tempNum
// 	tempNum = tempNum / 10
// }
// fmt.Println()
// }
// func main(){
// orderNum(1234535)
// }

// 12 exercise.

// package main

// import "fmt"

// func tubBoluvchi(N int) {
// 	num := N
// 	counter := 0package main

// import(
// 	"fmt"
// 	"math/rand"
// )
// func main(){
// 	slice := make([]int,5)
// 	for i := 0; i < 5; i++ {
// 		slice[i] = rand.Intn(10)
// 	}
// 	fmt.Println()
// }
// 	limit := num + 1
// 	for i := num; i < limit; i++ {
// 		for j := 2; j < i/2; j++ {
// 			counter = 0
// 			if i%j == 0 {
// 				for g := 1; g < j; g++ {
// 					if j%g == 0 {
// 						counter++
// 					}
// 				}
// 				if counter < 2 {
// 					fmt.Print(j, " ")
// 				}

// 			}
// 		}
// 	}
// 	fmt.Println()
// }

// func main() {
// 	tubBoluvchi(25)
// }

// 13 exercise.

// package main

// import "fmt"

// func multSon(N int){
// tempNum := N
// sum := 1
// for tempNum != 0 {
// 	sum = sum * tempNum
// 	tempNum = tempNum - 1
// }
// fmt.Println(sum)
// }

// func main(){
// multSon(5)
// }

// 14 exercise.

// package main

// import "fmt"

// func threeNum(n1,n2,n3 int){
// a,b,c := n1,n2,n3
// if a + b == c {
// 	fmt.Println(1)
// }else {
// 	fmt.Println(0)
// }
// }

// func main(){
// threeNum(1,2,3)
// }

// 15 exercise.

// package main

// import "fmt"

// func degreeNum(n1, n2 float32){
// num := n1
// d :=  n2
// var sum float32 = 1
// var i float32
// for i = 0; i < d;i++ {
// sum = sum * num
// }
// fmt.Println(sum)
// }

// func main(){
// degreeNum(0.5, 4)
// }

// 16 exercise.

// package main

// import "fmt"

// func toqSon(num int) {
// len := num
// for i := 0; i < len;i++{
// 	if i % 2 != 0 {
// 		fmt.Print(i, " ")
// 	}
// }
// fmt.Println()
// }

// func main(){
// 	toqSon(6)
// }

// 17 exercise.

// package main

// import "fmt"

// func multiplyNum(){
// 	var N int
// 	fmt.Print("enter N: ")
// 	fmt.Scan(&N)
// 	len := N
// 	sum := 0
// 	for i := 1; i <= len; i++ {
// 		for j := 1; j < 11; j++ {
// 			sum = i * j
// 			fmt.Println(i,"*",j,"=",sum)
// 		}
// 		fmt.Println()
// 	}
// }

// func main(){
// multiplyNum()
// }

 