package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "os"
)

func CreateConnection() (*gorm.DB, error) {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    //dbName := os.Getenv("DB_NAME")
    dbName := "blue"
    password := os.Getenv("DB_PASSWORD")

    return gorm.Open(
        "mysql",
        fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
            user, password, host, dbName),
    )
}
