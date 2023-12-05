package main

// --PRACTICE--

// 1 exercise.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr1 = [...]int{1, -2, 3, -4, 5, -6, 7, -8, 9, -10}
// 	n := len(arr1)
// 	find := 0
// 	sum := 0
// 	var a, b int
// 	for i := 0; i < n; i++ {
// 		if arr1[i] < 0 {
// 			find++
// 			if find == 2 {
// 				a = i
// 			}
// 			if find == 4 {
// 				b = i
// 			}
// 		}
// 	}

// 	for i := a+1; i < b; i++ {
// 		sum += arr1[i]
// 	}

// 	fmt.Println(sum)
// }

// 2 exercise.

// package main

// import "fmt"

// func main() {
// 	arrson := [...]int{0, -1, 2, -3, 4, 0, 5, -6, 7, 0}
// 	someText := len(arrson)
// 	musbat := 0
// 	manfiy := 0
// 	nollar := 0
// 	juftlar := 0
// 	toqlar := 0
// 	for i := 0; i < someText; i++ {
// 		if arrson[i] > 0 {
// 			musbat++
// 		}
// 		if arrson[i] < 0 {
// 			manfiy++
// 		}
// 		if arrson[i] == 0 {
// 			nollar++
// 		}
// 		if arrson[i]%2 == 0 {
// 			juftlar++
// 		}
// 		if arrson[i]%2 != 0 {
// 			toqlar++
// 		}
// 	}
// 	fmt.Println("Musbatlar: ", musbat, "ta")
// 	fmt.Println("Manfiylar: ", manfiy, "ta")
// 	fmt.Println("Nollar: ", nollar, "ta")
// 	fmt.Println("Juftlar: ", juftlar, "ta")
// 	fmt.Println("Toqlar: ", toqlar, "ta")
// }

// 3 exercise.

// package main

// import "fmt"

// func main() {
// 	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	someText := len(slice)
// 	fmt.Println("juftlar: ")

// 	for i := 0; i < someText; i++ {
// 		if slice[i]%2 == 0 {
// 			fmt.Print(i, " ")
// 		}
// 	}
// 	fmt.Println("\ntoq: ")
// 	for i := 0; i < someText; i++ {
// 		if slice[i]%2 != 0 {
// 			fmt.Print(i, " ")
// 		}
// 	}
// 	fmt.Print()
// }

// 4 exercise.

// package main

// import "fmt"

// func indexSum(){
// 	slice := []int{1,2,3,4,5,6,7,8,9,10}
// 	length := len(slice)
// 	sum1 := 1
// 	sum2 := 0
// 	for i := 0; i < length; i++ {
// 		if i != 0{
// 		if i % 2 == 0 {
// 		 sum1 = sum1 * slice[i]
// 		}
// 	}
// 	if slice[i] % 2 != 0 {
// 		sum2 = sum2 + slice[i]
// 	}
// }
// fmt.Println(sum1)
// fmt.Println(sum2)
// }

// func main(){
// indexSum()
// }

// 5 exercise.

// package main

// import "fmt"

// func minTwo(){
// slice := []int{1,2,3,4,5,6,7,8,9,10}
// length := len(slice)
// min := slice[0]
// secondmin := slice[1]

// for i := 1; i < length; i++ {
//  if slice[i] < min {
// 	secondmin = min
// 	min = slice[i]
//  }else if slice[i] < secondmin && slice[i] != min{
//   secondmin = slice[i]
//  }
// }
// fmt.Println("secondmin is:",secondmin)
// }

// func main(){
// 	minTwo()
// }

// 6 exercise.

// package main

// import "fmt"

// func maxTwo(){
// 	slice := []int{1,2,3,4,5,6,7,8,9,10}
// 	length := len(slice)
// 	max := slice[0]
// 	secondmax := slice[0]
// 	for i := 1; i<length; i++{
// 		if slice[i] > max{
// 			secondmax = max
// 			max = slice[i]
// 		}else if secondmax < slice[i] && slice[i] != max{
// 			secondmax = slice[i]
// 		}
// 	}
// 	fmt.Println("secondmax is",secondmax)
// }

// func main(){
// maxTwo()
// }

// 7 exercise.

// package main

// import (
// 	"fmt"
// )

// func changeNum(){
// slice := []int {1,2,3,4,5}
// length := len(slice)-1
// slice[0],slice[length] = slice[length],slice[0]
// fmt.Println(slice)
// }

// func main(){
// changeNum()
// }

// 8 exercise.

// package main

// import "fmt"

// func maxNum(){
// slice := []int{2,9,3,-4,5}
// length := len(slice)
// max := slice[0]
// counter := 0

// for i := 1; i < length; i++ {
// 	if max < slice[i] {
// 		max = slice[i]
// 	}
// }
// for i := 0; i < length; i++{
// 	if max == slice[i]{
// 		for j := i; j < length-1; j++ {
// 			counter++
// 		}
// }
// }
//  fmt.Println("maxdan keyin",counter,"ta son bor")
// }
// func main(){
// maxNum()
// }

// 9 exercise.

// package main

// import "fmt"

// func minNum(){
// 	slice := []int{2,9,3,-4,5}
// 	length := len(slice)
// 	min := slice[0]
// 	counter := 0

// 	for i := 1; i < length; i++{
// 		if slice[i] < min{
// 			min = slice[i]
// 		}
// 	}
// for i := 0; i < length; i++{
// 	if slice[i] == min{
// 		for j := i; j > 0; j--{
// 			counter++
// 		}
// 	}
// 	}
// 	fmt.Println("mindan oldin",counter,"ta son bor")
// }

// func main(){
// minNum()
// }

// 10 exercise.

// package main

// import "fmt"

// func maxmin() {
// 	slice := [5]int{2,9,3,-4, 5}
// 	length := len(slice)
// 	length2 := len(slice)
// 	min := slice[0]
// 	max := slice[0]
// 	counter := 0

// 	for i := 1; i < length; i++ {
// 		if max < slice[i] {
// 			max = slice[i]
// 		}
// 	}
// 	for i := 1; i < length; i++ {
// 		if slice[i] < min {
// 			min = slice[i]
// 		}
// 	}

// 	for i := 0; i < length; i++ {
// 		if slice[i] == max {
// 			for j := i; j < length2; j++ {
// 				if slice[j-1] == min {
// 					length2 = j - 1
// 				}
// 			}
// 		}
// 	}
// 	for i := 0; i < length; i++ {
// 		if slice[i] == max {
// 			for j := 2; j < length2; j++ {
// 				counter++
// 			}
// 		}
// 	}
// 	fmt.Println("max va min orasida", counter, "ta son bor")
// }

// func main() {
// 	maxmin()
// }

// --HOMEWORK--

// 1 exercise.

// package main

// import "fmt"

// func changeNum() {
//  n := 5
//  a := []int{2, 9, 3, -4, 5}

//  max := a[0]
//  min := a[0]
//  for i := 1; i < n; i++ {
//   if a[i] > max {
//    max = a[i]
//   }
//   if a[i] < min {
//    min = a[i]
//   }
//  }

//  maxIndex := 0
//  minIndex := 0
//  for i := 0; i < n; i++ {
//   if a[i] == max {
//    maxIndex = i
//   }
//   if a[i] == min {
//    minIndex = i
//   }
//  }

//  a[maxIndex], a[minIndex] = a[minIndex], a[maxIndex]

//  fmt.Println("Yangi", a)
// }

// func main(){
// 	changeNum()
// }


// 2 exercise.
 
// package main

// import "fmt"

// func prime() {
//  numbers := []int{7, 12, 31, 5, 19, 4, 13, 6, 11, 9}
//  primes := []int{}
//  primeCount := 0

//  isPrime := func(n int) bool {
//   if n <= 1 {
//    return false
//   }
//   for i := 2; i*i <= n; i++ {
//    if n%i == 0 {
//     return false
//    }
//   }
//   return true
//  }

//  fmt.Println("Tub sonlar:")
//  i := 0
//  for i < len(numbers) {
//   if isPrime(numbers[i]) {
//    primes = append(primes, numbers[i])
//    primeCount++
//    fmt.Println(numbers[i])
//   }
//   i++
//  }
//  fmt.Println("Tub sonlar soni:", primeCount)
// }

// func main(){
// 	prime()
// }

// package main

// import "fmt"

// func evenodd() {
//     input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//     even := []int{}
//     odd := []int{}

//     for i := 0; i < len(input); i++ {
//         num := input[i]
//         if num%2 == 0 {
//             even = append(even, num)
//         } else {
//             odd = append(odd, num)
//         }
//     }

//     fmt.Println("Juft:", even)
//     fmt.Println("Toq:", odd)
// }

// func main(){
// 	evenodd()
// }

// 3 exercise.
 
// package main

// import "fmt"

// func uniqueSon() {
//  var n int
//  fmt.Print("Massiv uzunligini kiriting: ")
//  fmt.Scan(&n)

//  arr := make([]int, n)
//  fmt.Println("Massiv elementlarini kiriting:")
//  for i := 0; i < n; i++ {
//   fmt.Scan(&arr[i])
//  }

//  indexSum := 0
//  valueSum := 0
//  uniqueValues := []int{}

//  for i := 0; i < n; i++ {
//   indexSum += i
//   valueSum += arr[i]

 
//   isUnique := true
//   for j := 0; j < len(uniqueValues); j++ {
//    if arr[i] == uniqueValues[j] {
//     isUnique = false
//     break
//    }
//   }
//   if isUnique {
//    uniqueValues = append(uniqueValues, arr[i])
//   }
//  }

//  if indexSum > valueSum {
//   fmt.Println("INDEX")
//  } else if valueSum > indexSum && len(arr) == len(uniqueValues) {
//   fmt.Println("WOW")
//  }
// }

// func main(){
// 	uniqueSon()
// }

// 4 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
// )

// func randNum () {
//  rand.Seed(rand.Int63()) 

//  var n int
//  fmt.Print("Massiv uzunligini kiriting: ")
//  fmt.Scan(&n)

//  arr := make([]int, n)
//  for i := 0; i < n; i++ {
//   arr[i] = rand.Intn(39) - 12  
//  }

//  fmt.Println("Input:", arr)

//  for i := 0; i < n; i++ {
//   if arr[i] < 0 {
//    arr[i] = 0
//   } else {
//    arr[i] = 1
//   }
//  }

//  fmt.Println("Output:", arr)
// }
// func main(){
// 	randNum()
// }


// 5 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
//  "time"
// )

// func randNum() {
//  rand.Seed(time.Now().UnixNano())

//  var n int
//  fmt.Print("Massiv uzunligini kiriting: ")
//  fmt.Scan(&n)

//  arr := make([]int, n)
//  for i := 0; i < n; i++ {
//   arr[i] = rand.Intn(22) + 14  
//  }
//  fmt.Println("Input:", arr)

//  for i := 0; i < n; i++ {
//   if arr[i]%2 == 0 {
//    fmt.Print(arr[i], " + ")
//   }
//  }

//  fmt.Println()
// }
// func main(){
// 	randNum()
// }


// 6 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
//  "time"
// )

// func randNum() {
//  rand.Seed(time.Now().UnixNano())

//  var n int
//  fmt.Print("Massiv uzunligini kiriting: ")
//  fmt.Scan(&n)

//  arr := make([]int, n)
//  for i := 0; i < n; i++ {
//   arr[i] = rand.Intn(52) - 46 

//  fmt.Println("Input:", arr)

//  for i := 0; i < n; i++ {
//   if arr[i]%2 != 0 {
//    fmt.Print("(", arr[i], ")")
//   }
//  }

//  fmt.Println()
// }
// }
// func main(){
// 	randNum()
// }

 
// 7 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
//  "time"
// )

// func randNum() {
//  rand.Seed(time.Now().UnixNano())

//  var n int
//  fmt.Print("Massiv uzunligini kiriting: ")
//  fmt.Scan(&n)

//  arr := make([]int, n)
//  for i := 0; i < n; i++ {
//   arr[i] = rand.Intn(31) - 5  
//  }

//  fmt.Println("Input:", arr)

//  shiftedArr := make([]int, n)
//  for i := 0; i < n-1; i++ {
//   shiftedArr[i] = arr[i+1]
//  }
//  shiftedArr[n-1] = arr[0]

//  fmt.Println("Output:", shiftedArr)
// }

// func main(){
// 	randNum()
// }


// 8 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
//  "time"
// )

// func randNum() {
//  n := 5
//  array := []int{10, 7, 0, -5, -2}

//  rand.Seed(time.Now().UnixNano())
//  randomSlice := make([]int, n)
//  perm := rand.Perm(len(array))
//  for i := 0; i < n; i++ {
//   randomSlice[i] = array[perm[i]]
//  }

//  result := append(randomSlice, array...)
//  fmt.Println(result)
// }

// func main(){
// 	randNum()
// }


// 9 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
//  "time"
// )

// func randNum() {
//  n := 5
//  array := []int{-10, 5, 0, 1, -2}

//  rand.Seed(time.Now().UnixNano())
//  randomSlice := make([]int, n)
//  for i := 0; i < n; i++ {
//   randomSlice[i] = rand.Intn(24) - 15
//  }

//  result := append(randomSlice, array...)
//  for i := len(result) - 1; i > 0; i-- {
//   result[i], result[i-1] = result[i-1], result[i]
//  }

//  fmt.Println(result)
// }

// func main(){
// 	randNum()
// }


// 10 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
//  "time"
// )

// func randNum() {
//  n := 10
//  array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//  k := 3

//  rand.Seed(time.Now().UnixNano())
//  randomSlice := make([]int, n)
//  for i := 0; i < n; i++ {
//   randomSlice[i] = rand.Intn(14) + 8
//  }

//  result := append(randomSlice, array...)
//  result = append(result[k:], result[:k]...)

//  fmt.Println(result)
// }
// func main(){
// 	randNum()
// }


// 11 exercise.

// package main

// import (
//  "fmt"
//  "math/rand"
//  "time"
// )

// func newarray() {
//  n := 10
//  min := -11
//  max := 70
//  k := 3

 
//  array := make([]int, n)
//  rand.Seed(time.Now().UnixNano())
//  for i := 0; i < n; i++ {
//   array[i] = rand.Intn(max-min+1) + min
//  }
//  fmt.Println("Original array:", array)

//  k = k % n
//  newArray := make([]int, n)
//  copy(newArray, array[n-k:])
//  copy(newArray[k:], array[:n-k])
//  fmt.Println("yangi array:", newArray)
// }
// func main (){
// 	newarray()
// }