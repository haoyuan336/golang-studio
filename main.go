package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

type UserData struct {
	Name string
	Text string
}

func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	t, err := template.New(file).ParseFiles("view/" + file)
	checkErr(err)
	t.Execute(w, data)
}

func writeData(userdata *UserData) string {
	///打开数据库
	db, err := sql.Open("sqlite3", "./data/db")
	checkErr(err)
	defer db.Close()
	//如果数据表不存在则创建
	db.Exec("create table data (id integer not null primary key,name text,data string);")
	var olddata string //数据库中已存在的数据
	var sqlStmt string //sql内容
	//查询用户是否存在，同事读取用户数据
	// err = db.QueryRow("select data from data where name = ?",userdata.Name).Scan(%olddata)

}
func index(w http.ResponseWriter, r *http.Request) {
	// renderHTML(w, "index.html", "no data")
	t, err := template.New("index.html").ParseFiles("view/" + "index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "no data")
}
func page(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Println("Handle: page: ParseForm:", err)
		}
		//获取客户端输入的内容
		u := UserData{}
		u.Name = r.Form.Get("username")
		u.Text = r.Form.Get("usertext")
		renderHTML(w, "page.html", u)
	} else {
		///如果不是通过post提交的数据.则将页面重定向到主页
		renderHTML(w, "redirect.html", "/")
	}
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/page", page)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
