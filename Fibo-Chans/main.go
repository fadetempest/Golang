package main

import (
	"fmt"
	"time"
)

func fib(c chan int){
	a := <- c
	arr:=make([]int, a)
	arr[0] = 0
	arr[1] = 1
	for i:=2; i < len(arr); i++{
		arr[i] = arr[i-1]+arr[i-2]
	}
	c <- arr[a-1]
}


func main()  {
	var c chan int=make(chan int, 2)
	var num int

	fmt.Scan(&num)
	c <- num

	go fib(c)
	time.Sleep(time.Millisecond)
	fmt.Println(<-c)
}