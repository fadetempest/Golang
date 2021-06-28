package main

import (
	"fmt"
)

func main(){
	for value:= range work(){
		fmt.Print(string(value))
	}
	/*var text rune
	wg:= sync.WaitGroup{}
	fmt.Scan(&text)
	wg.Add(1)
	go convert(text)
	wg.Wait()*/
}

func work() chan int{
	ch:=make(chan int)
	go convert(ch)
	return ch
}

func convert(ch chan int){
	defer close(ch)
	var text string
	fmt.Scan(&text)
	for _, letter:= range text{
		if string(letter) == " "{
			ch <- int(32)
		} else if rune(letter) != 121 && rune(letter) != 122{
			ch <- int(rune(letter))+ 2
		} else if rune(letter) == 121{
			ch <- int(97)
		} else if rune(letter) == 122{
			ch <- int(98)
		}
	}
}

