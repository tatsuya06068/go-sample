package main

import ( 
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    fmt.Println("Hello, World!")

    dbconf := "webuser:webpass@tcp(0.0.0.0:3306)/go_mysql8_development?charset=utf8mb4"

    db, err := sql.Open("mysql", dbconf)
    // 接続が終了したらクローズする
    defer db.Close()  
    fmt.Println(err);
    
    if err != nil {
        fmt.Println(err.Error())
    }


    

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }

}