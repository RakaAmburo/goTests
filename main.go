package main

import (
    "fmt"
)

func main(){
    fmt.Println("start")

    writer := CsvWriter{}
    writer.Init("test1.txt")
}
