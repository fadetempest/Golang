package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
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
	wg:= sync.WaitGroup{}
	fmt.Print("Введите текст\n")
	text := Scan()
	var codedText string
	for value:= range working(text){
		codedText += string(value)
	}
	wg.Add(2)
	go codToFile(codedText,wg)
	go nonCodToFile(text,wg)
	wg.Wait()
}

func working(text string) chan int{
	ch:=make(chan int)
	go converting(ch, text)
	return ch
}

func converting(ch chan int, text string){
	defer close(ch)
	var point int
	fmt.Print("Введите сдвиг: ")
	fmt.Scan(&point)
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

func codToFile(codedText string, wg sync.WaitGroup) {
	content := []byte(codedText)
	err := ioutil.WriteFile("codedText.txt", content, 0644)
	if err != nil {
		panic(err)
	}
	wg.Done()
}

func nonCodToFile(nonCodText string, wg sync.WaitGroup){
	content := []byte(nonCodText)
	err := ioutil.WriteFile("nonCodedText.txt", content, 0644)
	if err != nil {
		panic(err)
	}
	wg.Done()
}