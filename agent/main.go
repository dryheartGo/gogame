package main

import (
    "fmt"
    "log"
    "os"
    "utils/strs"
)

func main() {
    fmt.Println(strings.IsBlank("222"))
    var s string
    //var b, f, s = true, 2.3, "four"

    name := "/filename"
    _, err := os.Open(name)
    if err != nil {
        log.Println(err)
    }

    fmt.Printf(s)

    //指针
    x := 1
    p := &x
    fmt.Println(*p) // "1"
    *p = 2
    fmt.Println(x) // "2"

    //数组
    var a [3]int
    for i, v := range a {
        //0 0
        fmt.Printf("%d %d\n", i, v) //1 0
    } //2 0

    //切片
    i := []int{0, 1, 2, 3, 4, 5}
    str := make([]string, 10, 20)
    fmt.Println(i, str) //[0 1 2 3 4 5] [         ]
    i1 := i[2:4]
    i2 := i[:5]
    fmt.Println(i1, i2) //[2 3] [0 1 2 3 4]
}
