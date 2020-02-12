package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("mysql", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s","127.0.0.1","proxysql",""))
	if db == nil || err != nil {
		return
	}
	defer db.Close()

}
