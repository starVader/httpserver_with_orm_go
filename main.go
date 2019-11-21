package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func (Employee) TableName() string {
	return "employee"
}
func dbConn() *gorm.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "smokesow#123"
	dbName := "goblog"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var employee = Employee{}
	res := db.Find(&employee)
	defer db.Close()
	data, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(res)
	_, _ = w.Write(data)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
