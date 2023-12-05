package main

import "fmt"

func main(){
	p := Programmer{
		job: "programmer",
	}
	t := Teacher{
		job: "teacher",
	}

	printJob(p)
	printJob(t)
}

type Worker interface{
Work()
}

func printJob (w Worker){
	w.Work()
}

type Programmer struct{
job string
}

func(p Programmer) Work(){
	fmt.Println("Job of first worker is:",p.job)
}

type Teacher struct{
job string
}

func (t Teacher) Work(){
	fmt.Println("Job of second worker is:", t.job)
}

