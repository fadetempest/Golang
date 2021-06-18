package main

import "fmt"

func main(){
	var a int
	var b int
	fmt.Println("Введите первое число")
	fmt.Scan(&a)
	fmt.Println("Введите второе число")
	fmt.Scan(&b)
	ar :=make([]int, b) //Create an empty array
	ar[0] = 0 //First Fibonacci number
	ar[1] = 1 //Second Fibonacci number
	for i:=2; i < len(ar); i++{
		ar[i] = ar[i-1] + ar[i-2] //Fill array with Fibonacci numbers
	}
	for i:=a - 1; i < b; i++{
		fmt.Print(ar[i], " ") //Printing numbers from a to b
	}
}
