package main

import (
	"fmt"
	"time"
)

func main() {
	cha3()
}


/*
> cha0()
0
*/
func cha0() {
	ch := 0

	go func() {
		time.Sleep(time.Second)	
		ch = 10
	}()

	fmt.Println(ch)
}


/*
> cha1()
10
*/
func cha1() {
	ch := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)	
		ch <- 10
	}()

	fmt.Println(<-ch)
}


/*
> cha2()
30
*/
func hoge(ch chan int) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		ch <- 10
	}
}

func cha2() {
	ch := make(chan int, 3)

	go hoge(ch)

	fmt.Println(<-ch + <-ch + <-ch)
}


/*
> cha3()
40
*/
func foo(ch chan int) {
	time.Sleep(time.Second)
	ch <- 10
}

func cha3() {
	ch := make(chan int, 3)

	go foo(ch)

	go foo(ch)

	go foo(ch)
	x := <-ch + <-ch + <-ch

	foo(ch)
	x += <-ch
	fmt.Println(x)
}