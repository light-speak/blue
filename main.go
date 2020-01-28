package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "")
	if db == nil || err != nil {
		return
	}
	defer db.Close()

}
