package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//gor0()
	//gor1()
	//gor2()
	gor3()
}

var wg sync.WaitGroup


/*
> gor0()
hoge
foo

or 

hoge
*/
func gor0() {
	go fmt.Println("foo")
	fmt.Println("hoge")
}


/*
> gor1()
hoge: 0
foo: 0
hoge: 1
foo: 1
hoge: 2
*/
func gor1() {
	go func () {
		for i := 0; i < 3; i++ {
			fmt.Println("foo:", i)
			time.Sleep(2 * time.Second)
		}
	}()

	func () {
		for i := 0; i < 3; i++ {
			fmt.Println("hoge:", i)
			time.Sleep(time.Second)
		}
	}()
}


/*
> gor2()
hoge: 0
foo: 0
hoge: 1
foo: 1
hoge: 2
foo: 2
*/

func gor2() {
	//goルーチンのカウント +1
	wg.Add(1)
	go func () {
		//goルーチンのカウント -1
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("foo:", i)
			time.Sleep(2 * time.Second)
		}
	}()

	func () {
		for i := 0; i < 3; i++ {
			fmt.Println("hoge:", i)
			time.Sleep(time.Second)
		}
	}()

	//goルーチンのカウントが0になるまで待つ
	wg.Wait()
}


func foo() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("foo:", i)
		time.Sleep(2 * time.Second)
	}
}

func hoge() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("hoge:", i)
		time.Sleep(time.Second)
	}
}

/*
> gor3()
foo: 0
hoge: 0
hoge: 0
hoge: 1
hoge: 1
hoge: 2
hoge: 2
foo: 1
*/
func gor3() {
	wg.Add(2)
	go hoge()
	go hoge()
	go foo()
	wg.Wait()
}