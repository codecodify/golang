package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func handleIndex(w http.ResponseWriter, r *http.Request){
	var message string
	var status string

	// 判断是否post请求
	if(r.Method == "POST"){
		// 接收数据

		// 解析传递的参数，如果用r.FormValue()，则可省略
		err := r.ParseForm()
		if err != nil{
			fmt.Println(err)
		}
		username := strings.Join(r.Form["username"], "")
		password := strings.Join(r.Form["password"], "")
		if username == "admin" && password == "admin"{
			message = "登陆成功"
			status = "success"
		}else{
			message = "账号或密码错误"
			status = "danger"
		}
	}

	// 引入模板
	t := template.Must(template.ParseFiles("template/index.html"))

	// 传入模板的数据
	data := map[string]string{
		"message": message,
		"status": status,
	}
	t.Execute(w, data)
}
