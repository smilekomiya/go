package main

import(
    "fmt"
)

func Sqrt(x float64) float64 {
    var z float64 = 1
    var i int = 1
    for ; i <= 10000000; i++ {
        z = z - (z * z - x) / 2 * z
    }

    return z
}

func main() {
    fmt.Println(Sqrt(2))
}