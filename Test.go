package main;

import (
//"encoding/json"
"net/http"
"fmt"
"time"
"os"
)

type Message struct{
	Body string
	Time string
}

var spisoc [1000]Message
var Strok string
var index int = 0

//Функции - обработкичи.

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, это мой сервак, он работает!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Московское время: %s\n", time.Now().String())
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	//Место, где расположены файлы (на самом деле в /go/src/github.com/user/Test/static).
	base := os.Getenv("HOME") + "/go/src/github.com/user/Test"
	
	//Теперь нам надо изьять название файла из URL-ссылки.
	file := r.RequestURI
	
	
	
	// Если не удалось загрузить нужный файл – показать сообщение о  ошибке.
	reqFile, err := os.Open(base + file)
	
	if err == nil {
	//Считать содержимое файла.
	fi, _ := reqFile.Stat()
	bytes := make([]uint8, fi.Size())
	reqFile.Read(bytes)	
	//Отобразим содержимое на браузере.
	fmt.Fprintf(w,"%s\n", string(bytes))
	}
	
	if err != nil {
		fmt.Fprintf(w, "Файл не найден!\n")
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Если был POST запрос, то выполняется этот код.
	if r.Method == "POST"{
		fmt.Fprintf(w,"Вас запрос принят:\n")
		r.ParseForm()//Вытаскиваем данные из запроса.
		//fmt.Fprintf(w,"Он принял: %s\n", r.PostForm)
		// Теперь нам надо запрос сделать нормальным, а не в виде среза.
		for key, _ := range r.PostForm {
			fmt.Fprintf(w,"Он принял: %s\n", key)
			spisoc[index].Body = key
		}
		spisoc[index].Time = time.Now().String()
		index++
	}
	// Если GET запрос - то этот.
	if r.Method == "GET"{
		fmt.Fprintf(w,"Список сообщений:\n")
		for i:=0; i<=index; i++{
			fmt.Fprintf(w,"%s  %s\n", spisoc[i].Body, spisoc[i].Time)
		}
		
	}
}

func clearHandler(w http.ResponseWriter, r *http.Request) {		
	for i:=0; i<=index; i++{
			spisoc[i].Body = ""
			spisoc[i].Time = ""
		}
}

func main(){
	
	// В этом блоке мы определяем функции обработкичики запросов
	http.HandleFunc("/", requestHandler) //Тут проверяется рабспособность сервера.
	http.HandleFunc("/time", timeHandler)//Здесь сервер выдает текущее время(СВОЕ).
	http.HandleFunc("/static/", staticHandler)//Отдает статические файлы из каталога на диске или страницу с ошибкой, если такого файла нет.
	http.HandleFunc("/message", messageHandler)//Добавляет новое сообщение в список (POST) и возвращает список всех сообщений (GET) в формате JSON
	http.HandleFunc("/message/clear", clearHandler)
	http.ListenAndServe(":8080", nil)//Тут задаем порт,на который передаем свои запросы.
}