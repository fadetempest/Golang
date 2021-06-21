package main

import "fmt"

func main(){
	for value := range worker(){
		fmt.Print(value, " ")
		}
	}


func worker() chan int {
	ch := make(chan int)
	go fib(ch)
	return ch
}

func fib(ch chan int) chan int {
	defer close(ch)
	a := 1
	b := 5
	arr:=make([]int, b)
	arr[0] = 0
	arr[1] = 1
	for i:=2;i<len(arr);i++{
		arr[i] = arr[i-1] + arr[i-2]
	}
	for i,item:= range arr{
		if (a - 1) <= i && i <= b{
			ch <- item
		}
	}
	return ch
}
