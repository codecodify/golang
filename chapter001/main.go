package main

import (
	"net/http"
	"html/template"
)

func main() {
	// 路由定义
	http.HandleFunc("/",  handleHome)
	// 开启服务，端口为8080
	http.ListenAndServe(":8080", nil)
}


func handleHome(w http.ResponseWriter, r *http.Request){
	// 引入模板文件
	t, _ := template.ParseFiles("./template/index.html")
	// 解析模板文件并输出内容
	t.Execute(w, nil)
}