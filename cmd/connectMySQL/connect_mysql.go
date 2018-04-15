package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"expvar"
)

//var db = &sql.DB{}


func main() {
	//user:password@/dbname

	//root:kLoBXVc7g@tcp(127.0.0.1:3306)/Shop?charset=utf8

	query(db)

	//insert(db)
}

func query(db *sql.DB)  {
	rows, err := db.Query("SELECT  `name`, age FROM shop_user")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan( &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" %s is %d\n", name, age)
	}
}




func insert(db *sql.DB)  {
	stmt, err := db.Prepare("insert into shop_user(name, age) values(?,?)")
	defer  stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}

	stmt.Exec("lee", 32)
	stmt.Exec("wen", 50)
}

func update(db *sql.DB)  {
}

//func init()  {
//	_, err := sql.Open("mysql", "root:kLoBXVc7g@mysql")
//}