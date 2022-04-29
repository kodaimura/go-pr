package main

import (
	"fmt"
)

func main() {
	gen1()
}

type Number interface {
    int | int32 | int64 | float32 | float64
}

type String interface {
    string
}


/*
> gen0()
3
4.6
ab
*/
func gen0() {
	fmt.Println(plus(1, 2))
	fmt.Println(plus(1.2, 3.4))
	fmt.Println(plus("a", "b"))
}

func plus[T Number | String](x, y T) T {
	return x + y
}


/*
> gen1()
[[c d] [c b] [d b] [b a] [c a] [d a]]
[[3 4] [3 2] [4 2] [2 1] [3 1] [4 1]]
*/
func gen1() {
    sl := []string{"a", "b", "c", "d"}
    fmt.Println(Combinations(sl, 2))

    sl2 := []int{1, 2, 3, 4}
    fmt.Println(Combinations(sl2, 2))
}


func Combinations[T any](sl []T, n int) [][]T {
    combs := [][]T{}

    if len(sl) <= n{
        return [][]T{sl} 
    } 

    if n == 1 {
        for _, x := range sl {
            combs = append(combs, []T{x})
        }
        return combs
    }

    for _, c := range Combinations(sl[1:], n - 1) {
        combs = append(combs, append(c, sl[0]))
    }    
    return append(Combinations(sl[1:], n), combs...)
}