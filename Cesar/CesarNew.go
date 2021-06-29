package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	firstPoint = 65
	secondPoint = 91
	thirdPoint = 90
	fourPoint = 123
	ninetySeven = 97
	deleter = 26
	lastPoint = 122
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
		if nonChanging(letter){
			ch <- int(rune(letter))
		} else if changingLetter(letter, point){
			ch <- int(rune(letter))-deleter+point
		} else {
			ch <- int(rune(letter)) + point
		}
	}
}

func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil{
		fmt.Printf("Error while reading %s\n", err)
	}
	return str
}

func nonChanging(letter int32) bool{
	return (int(rune(letter)) < firstPoint) || ((int(rune(letter)) > thirdPoint) && (int(rune(letter)) < ninetySeven)) || (int(rune(letter)) > lastPoint)
}

func changingLetter(letter int32, point int) bool{
	return (int(rune(letter)) + point > secondPoint && (int(rune(letter)) + point < ninetySeven)) || (int(rune(letter)) + point > fourPoint)
}
