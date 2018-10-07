package main

import "fmt"

func main() {
    var appended = []int{1, 2, 3}
    var slice = Filter(appended, func(arg int) bool {
        return arg % 2 == 1
    })
    fmt.Println(slice)
}

func Filter(s []int, fn func(int) bool) []int {
    var p []int
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}