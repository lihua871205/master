package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
)

func helloHandler (w http.ResponseWriter, r * http.Request){
	fmt.Println("helloHandler")
	io.WriteString(w, "金殷荷是个大笨蛋\r\n坐着飞机丢炸弹\r\n炸死人民千千万")
}

func main(){
	fmt.Println("main begin!")
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("main begin!1")
	err := http.ListenAndServe("localhost:8000", nil)
	fmt.Println("main begin!2")

	if (err != nil){
		log.Fatal("ListenAndServer:", err.Error())
	}

	fmt.Println("success!")
}
