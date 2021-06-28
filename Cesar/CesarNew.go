package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	for value:= range working(){
		fmt.Print(string(value))
	}
}

func working() chan int{
	ch:=make(chan int)
	go converting(ch)
	return ch
}

func converting(ch chan int){
	defer close(ch)
	var point int
	fmt.Print("Введите сдвиг: ")
	fmt.Scan(&point)
	text := Scan()
	for _, letter:= range text{
		if (64 > int(rune(letter)) && int(rune(letter)) < 91) && (96 > int(rune(letter)) && int(rune(letter)) < 123){
			ch <- int(rune(letter))
		} else if (64 > int(rune(letter)) + point && int(rune(letter)) + point < 91) || (96 > int(rune(letter)) + point || int(rune(letter)) + point > 123){
			ch <- int(rune(letter))-26+point
		} else {
			ch <- int(rune(letter)) + point
		}
	}
}

func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	return str
}
