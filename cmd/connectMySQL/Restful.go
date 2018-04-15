package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	Seq int
	Result map[string] int
}

type Message struct {
	Code    string
	Msg string
	Time    int64
	UserInfo UserInfo
}

type UserInfo struct {
	Id int
	Name string
	Age int
}

//先定义一个全局变量
var db *sql.DB;
var err error;

func init()  {
	//初始化mysql 驱动 获得db
	//db, err = sql.Open("mysql","root:123456@tcp(localhost:3306)/shop?utf8mb4");
	db, err = sql.Open("mysql","root:xixi1234@tcp(66.98.114.41:3306)/shop?utf8");
	fmt.Println(db);
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	pingErr := db.Ping();
	if pingErr != nil {
		log.Fatalf("Error on opening database connection: %s", pingErr.Error())
	}
	checkErr(err)
}

func main() {
	//设置监听路径
	http.HandleFunc("/", handler)
	//http.ListenAndServe("66.98.114.41:80", nil)
	http.ListenAndServe("localhost:8899", nil)
}

func queryUserInfo(name string)([]byte, error) {
	//查询数据库
	rows, err := db.Query("select id, `name`, age from shop_user where name = ?", name);
	checkErr(err);
	var uInfo UserInfo;
	for rows.Next() {
		var id int
		var name string
		var age int
		if err:= rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err);
		}
		uInfo = UserInfo{id, name, age};
	}
	//拼接数据格式 并且是json 返回
	m := Message{"0","获取成功", time.Now().Unix(),uInfo}
	return json.MarshalIndent(m,"","");
}


func getJson()([]byte, error)  {
	pass := make(map[string] int)
	pass["x"] = 50
	pass["c"] = 60
	//item1 := Item{100, pass}

	reject := make(map[string] int)
	reject["l"] = 11
	reject["d"] = 20
	//item2 := Item{200, reject}

	u := UserInfo{1,"xixi",23}
	//detail := []Item{item1, item2}
	fmt.Println(u)
	//m := Message{"0", "获取成功", time.Now().Unix()}
	//fmt.Println(m)
	m := Message{"0", "获取成功", time.Now().Unix(), u}


	return json.MarshalIndent(m,"", "")
}

func handler(w http.ResponseWriter, r *http.Request)  {
	//resp, err := getJson();
	resp, err := queryUserInfo("xixi");
	checkErr(err);
	fmt.Fprintln(w, string(resp))
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}