package main

import (
    "fmt"
    "dbooks/src/entities"
    "dbooks/src/config"
)

func main (){
    db,err := config.GetMySQLDB()
    if err != nil {
        fmt.Println(err)
    }
    var book entities.Book
    book.Name="22"
    fmt.Printf(book.Name)
    fmt.Printf(book.Name)
}
