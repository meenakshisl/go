package main

import (
    "fmt"
    "math"
    "reflect"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n  //automatic typr conversion
    fmt.Println(reflect.TypeOf(d)) //returns the type of d

    fmt.Println(int64(d))
    fmt.Println(math.Sin(n)) //sin needs float64 type argument, int64 is automatically converted to float64
}
