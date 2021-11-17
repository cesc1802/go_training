package main

import (
	"database/sql"
	"fmt"
	mysql "github.com/cesc1802/go_training/internal/storages/mysql"
	"net/http"

	"github.com/cesc1802/go_training/internal/services"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//mysql
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/dataTest?charset=utf8")
	defer db.Close()
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer db.Close()
	http.ListenAndServe(":5050", &services.ToDoServiceMySQL{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &mysql.MySQLDB{
			DB: db,
		},
	})
}
