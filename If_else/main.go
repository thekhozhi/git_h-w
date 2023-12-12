package main

// --IF ELSE--
// 1 exercise.

// package main

// import "fmt"

// func main(){
// 	var A int
// 	var B int
// 	fmt.Println("enter A:")
// 	fmt.Scan(&A)
// 	fmt.Println("enter B:")
// 	fmt.Scan(&B)

// 	if A < B {
// 		fmt.Println(A,B)
// 	}else if A > B {
// 		A,B = B,A
// 		fmt.Println(A,B)
// 	}else{
// 		B++
// 		fmt.Println(A,B)
// 	}
// 	}

// 2 exercise.

// package main

// import "fmt"

// func main(){
// 	var(
// 		a int
// 		b int
// 		zero int = 0
// 	)
// 	fmt.Println("enter a:")
// 	fmt.Scan(&a)
// 	fmt.Println("enter b:")
// 	fmt.Scan(&b)

// 	if a != b {
// 		fmt.Println(a+b)
// 	}else {
// 		fmt.Println(zero)
// 	}
//}

// 3 exercise.

// package main

// import "fmt"

// func main(){
// 	var a int
// 	var b int
// 	var zero int = 0

// 	fmt.Println("enter a:")
// 	fmt.Scan(&a)
// 	fmt.Println("enter b:")
// 	fmt.Scan(&b)

// 	if a > b {
// 		fmt.Println(a)
// 	}else if a < b {
// 		fmt.Println(b)
// 	}else{
// 		fmt.Println(zero)
// 		fmt.Println(a,b)
// 	}
//}

// 4 exercise.

// package main

// import "fmt"

// func main(){
// 	var(
// 		a int
// 		b int
// 		c int
// 	)
// 	fmt.Println("enter a:")
// 	fmt.Scan(&a)
// 	fmt.Println("enter b:")
// 	fmt.Scan(&b)
// 	fmt.Println("enter c:")
// 	fmt.Scan(&c)

// 	if a < b && a < c {
// 		fmt.Println(a)
// 	}else if b < a && b < c {
// 		fmt.Println(b)
// 	}else {
// 		fmt.Println(c)
// 	}
// }

// 5 exercise.

// package main

// import "fmt"

// func main(){
// 	var(
// 		a int
// 		b int
// 		c int
// 	)
// 	fmt.Println("enter a:")
// 	fmt.Scan(&a)
// 	fmt.Println("enter b:")
// 	fmt.Scan(&b)
// 	fmt.Println("enter c:")
// 	fmt.Scan(&c)

// 	if a > b && a < c || a > c && a < b{
// 		fmt.Println(a)
// 	}else if b > a && b < c || b < a && b > c {
// 		fmt.Println(b)
// 	}else{
// 		fmt.Println(c)
// 	}
// }

// 6 exercise.

// package main

// import "fmt"

// func main(){
// var(
// 	a int
// 	b int
// 	c int
// )
// fmt.Print("enter a: ")
// fmt.Scan(&a)
// fmt.Print("enter b: ")
// fmt.Scan(&b)
// fmt.Print("enter c: ")
// fmt.Scan(&c)

// 	if a < b && a < c {
// 		fmt.Println(a)
// 	}else if b < a && b < c {
// 		fmt.Println(b)
// 	}else {
// 		fmt.Println(c)
// 	}

// 	if a > b && a > c {
// 		fmt.Println(a)
// 	}else if b > a && b > c {
// 		fmt.Println(b)
// 	}else {
// 		fmt.Println(c)
// 	}
// }

// 7 exercise.

// package main

// import "fmt"

// func main(){

// 	var(
// 		a int
// 		b int
// 		c int
// 	)
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&a)
// 	fmt.Print("enter b: ")
// 	fmt.Scan(&b)
// 	fmt.Print("enter c: ")
// 	fmt.Scan(&c)

// 	if a > b && a > c {
// 		fmt.Println(a)
// 	}else if b > a && b > c {
// 		fmt.Println(b)
// 	}else {
// 		fmt.Println(c)
// 	}

// 	if a > b && a < c || a > c && a < b{
// 		fmt.Println(a)
// 	}else if b > a && b < c || b < a && b > c {
// 		fmt.Println(b)
// 	}else{
// 		fmt.Println(c)
// 	}
// }

// 8 exercise.

// package main

// import "fmt"

// func main(){

// 	var(
// 		a int
// 		b int
// 		c int
// 	)
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&a)
// 	fmt.Print("enter b: ")
// 	fmt.Scan(&b)
// 	fmt.Print("enter c: ")
// 	fmt.Scan(&c)
// 		if a < b && b < c {
// 			a = a * 2
// 			b = b * 2
// 			c = c * 2
// 		fmt.Println(a,b,c)
// 		}else if a >= 0 && b >= 0 && c >= 0 {
// 			a = a - a - a
// 			b = b - b - b
// 			c = c - c - c
// 		fmt.Println(a,b,c)
// 		}else{
// 			a = a - a - a
// 			b = b - b - b
// 			c = c - c - c
// 		fmt.Println(a,b,c)
// 		}
// }

// 9 exercise.

// package main

// import "fmt"

// func main(){

// 	var(
// 		a int
// 		b int
// 		c int
// 	)
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&a)
// 	fmt.Print("enter b: ")
// 	fmt.Scan(&b)
// 	fmt.Print("enter c: ")
// 	fmt.Scan(&c)
// 		if a < b && b < c || a > b && b > c{
// 			a = a * 2
// 			b = b * 2
// 			c = c * 2
// 		fmt.Println(a,b,c)
// 		}else if a >= 0 && b >= 0 && c >= 0 {
// 			a = a - a - a
// 			b = b - b - b
// 			c = c - c - c
// 		fmt.Println(a,b,c)
// 		}else{
// 			a = a - a - a
// 			b = b - b - b
// 			c = c - c - c
// 		fmt.Println(a,b,c)
// 		}
// }

// 10 exercise.

// package main

// import "fmt"

// func main(){
// 	var (
// 		a int
// 		b int
// 		c int
// 	)
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&a)
// 	fmt.Print("enter b: ")
// 	fmt.Scan(&b)
// 	fmt.Print("enter c: ")
// 	fmt.Scan(&c)

// 	if b == c && b != a{
// 		a = 1
// 		fmt.Println(a)
// 	}else if a == c && a != b{
// 		b = 2
// 		fmt.Println(b)
// 	}else{
// 		c = 3
// 		fmt.Println(c)
// 	}
// }

// 11 exercise.

// package main

// import "fmt"

// func main(){
// 	var (
// 		a int
// 		b int
// 		c int
// 		d int
// 	)
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&a)
// 	fmt.Print("enter b: ")
// 	fmt.Scan(&b)
// 	fmt.Print("enter c: ")
// 	fmt.Scan(&c)
// 	fmt.Print("enter d: ")
// 	fmt.Scan(&d)

// 	if b == c && b == d{
// 		a = 1
// 		fmt.Println(a)
// 	}else if a == c && a == d{
// 		b = 2
// 		fmt.Println(b)
// 	}else if a == b && a == d{
// 		c = 3
// 		fmt.Println(c)
// 	}else {
// 		d = 4
// 		fmt.Println(d)
// 	}

//}

// ---FOR---
// 1 exercise.

// package main

// import "fmt"

// func main(){
// 	var i int
// 	fmt.Print("enter i: ")
// 	fmt.Scan(&i)
// 	for i=i; i>1; i--{
// 		if i % 2 == 0{
// 			fmt.Println(i)
// 		}
// 		}
// 	}

// 2 exercise.

// package main

// import "fmt"

// func main(){
// 	var i int
// 	fmt.Print("enter i: ")
// 	fmt.Scan(&i)
// 	for i=i; i>1; i--{
// 		if i % 2 == 1{
// 			fmt.Println(i)
// 		}
// 		}
// 	}

// 3 exercise.

// package main

// import "fmt"

// func main(){
// 	for i := 10; i < 100; i++ {
// 		if i % 2 != 0 && i % 3 != 0 && i % 5 != 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// 4 exercise.

// 5 exercise.

// package main

// import "fmt"

// func main(){
// 	fmt.Print(" enter n: ")
// 	var n int
// 	fmt.Scan(&n)
// 	var m int = n
// 	var sum int = 0

// 	for n > 0 {
// 	n = n / 10
// 	sum++
// 	}
// 	var nol int = 1
// 	for i := 1; i < sum; i++ {
// 	nol = nol * 10
// 	}
// 	var k int
// 	for i := 0; i < sum; i++ {
// 	k = (m / nol)
// 	fmt.Print(k, " ")
// 	m = m % nol
// 	nol = nol / 10
// 	}
// 	fmt.Println()
// }

// 6 exercise.

// package main

// import "fmt"

// func main(){
// 	var n = 0
// 	fmt.Print("enter n: ")
// 	fmt.Scan(&n)

// 	for i := 0; i < n; i++{
// 		for j := i; j < n-1;j++ {
// 			fmt.Print(" ")
// 		}
// 		for k := 0; k<=i; k++{
// 			if k == 0 || k == i {
// 				fmt.Print("* ")
// 			} else if i == n - 1{
// 				fmt.Print("* ")
// 			}else {
// 				fmt.Print("  ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// --SWITCH--
// 1 exercise.

// package main

// import "fmt"

// func main(){
// 	fmt.Println("Robot qaerga qarab turibdi \n s - shimol \n j - janub \n q - sharq \n g - garb")
// 	var holat string
// 	fmt.Scan(&holat)
// 	fmt.Println("Robotga komanda bering \n 0 - togriga yurishni davom ettir \n 1 - chapga buril \n 2 - onga buril ")
// 	var kom int
// 	fmt.Scan(&kom)

// 	if holat == "s" {
// 	  switch kom {
// 	  case 0:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Shimol")
// 	  case 1:
// 		fmt.Println("Komandadan keyingi hozirgi holat: G'arb")
// 	  case 2:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Sharq")
// 	  }

// 	} else if holat == "j" {
// 	  switch kom {
// 	  case 0:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Janub")
// 	  case 1:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Sharq")
// 	  case 2:
// 		fmt.Println("Komandadan keyingi  hozirgi holat: G'arb")
// 	  }

// 	} else if holat == "q" {
// 	  switch kom {
// 	  case 0:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Sharq")
// 	  case 1:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Shimol")
// 	  case 2:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Janub")
// 	  }

// 	} else if holat == "g" {
// 	  switch kom {
// 	  case 0:
// 		fmt.Println("Komandadan keyingi hozirgi holat: G'arb")
// 	  case 1:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Janub")
// 	  case 2:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Shimol")
// 	  }
// 	}
// }

// 2 exercise.

// package main

// import "fmt"

// func main(){
// 	fmt.Println("\n Lokatr qaysi tomonga qarap turipti \n s - shimol \n j - janub \n q - sharq \n g - garb")
// 	var holat string
// 	fmt.Scan(&holat)
// 	fmt.Println("Lokatrga birinchi komandani bering  \n 0 - o'nga buril \n 1 - chapga buril \n 2 - 180 buril")
// 	var kom1 int
// 	fmt.Scan(&kom1)
// 	var holatA string
// 	fmt.Println("Lokatrga ikkinchi komandani bering  \n 0 - o'nga buril \n 1 - chapga buril \n 2 - 180 buril")
// 	var kom2 int
// 	fmt.Scan(&kom2)

// 	if holat == "s" {
// 	  switch kom1 {
// 	  case 0:
// 		holatA = "q"
// 	  case 1:
// 		holatA = "g"
// 	  case 2:
// 		holatA = "j"
// 	  }

// 	} else if holat == "j" {
// 	  switch kom1 {
// 	  case 0:
// 		holatA = "g"
// 	  case 1:
// 		holatA = "q"
// 	  case 2:
// 		holatA = "s"
// 	  }
//

// 	} else if holat == "q" {
// 	  switch kom1 {
// 	  case 0:
// 		holatA = "j"
// 	  case 1:
// 		holatA = "s"
// 	  case 2:
// 		holatA = "g"
// 	  }
// 	} else if holat == "g" {
// 	  switch kom1 {
// 	  case 0:
// 		holatA = "s"
// 	  case 1:
// 		holatA = "j"
// 	  case 2:
// 		holatA = "q"
// 	  }

// 	}

// 	if holatA == "s" {
// 	  switch kom2 {
// 	  case 0:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Sharq")
// 	  case 1:
// 		fmt.Println("Komandadan keyingi hozirgi holat: G'arb")
// 	  case 2:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Janub")
// 	  }

// 	} else if holatA == "j" {
// 	  switch kom2 {
// 	  case 0:
// 		fmt.Println("Komandadan keyingi hozirgi holat: G'arb")
// 	  case 1:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Sharq")
// 	  case 2:
// 		fmt.Println("Komandadan keyingi hozirgi holat: Shimol")
// 	  }
// 	} else if holatA == "q" {
// 		switch kom2 {
// 		case 0:
// 		  fmt.Println("Komandadan keyingi hozirgi holat: Janub")
// 		case 1:
// 		  fmt.Println("Komandadan keyingi hozirgi holat: Shimol")
// 		case 2:
// 		  fmt.Println("Komandadan keyingi hozirgi holat: G'arb")
// 		}

// 	  } else if holatA == "g" {
// 		switch kom2 {
// 		case 0:
// 		  fmt.Println("Komandadan keyingi hozirgi holat: Shimol")
// 		case 1:
// 		  fmt.Println("Komandadan keyingi hozirgi holat: Janub")
// 		case 2:
// 		  fmt.Println("Komandadan keyingi hozirgi holat: Sharq")
// 		}
// 	}
// }
// khoji