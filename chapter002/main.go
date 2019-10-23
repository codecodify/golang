package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 注意大小写问题，如果要在模板调用数据，变量名首字母一定要大写
type Student struct {
	Name string
	Street string
}

func main() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request){
	// 函数的映射
	funcMap := template.FuncMap{"number": Number}
	// 创建新模板，并命名为index.html，并把函数注入到模板中（导入模板前需要先注入函数，否则会出现错误）
	t := template.New("index.html").Funcs(funcMap)
	// Must函数主要是检查模板是否出现问题
	t = template.Must(t.ParseFiles("./template/index.html"))
	data := []Student{
		{
			Name: "Anna Awesome",
			Street: "Broome Street",
		},
		{
			Name: "Debbie Dallas",
			Street: "Houston Street",
		},
		{
			Name: "John Doe",
			Street: "Madison Street",
		},
	}
	//t.Execute(w, data)
	err := t.ExecuteTemplate(w, "index.html", data)
	if err != nil{
		fmt.Println(err)
	}
}

// 编号 （自定义函数，在模板中使用）
func Number(number int) int {
	return number + 1
}