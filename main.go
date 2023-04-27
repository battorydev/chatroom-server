package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool{
		return true
	},
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Hello World from Go!")
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		fmt.Println(err)
		return
	}

	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil{
			fmt.Println(err)
			return
		}

		fmt.Println(string(msg))
		err = socket.WriteMessage(msgType, msg)
		if err != nil{
			fmt.Println(err)
			return
		}

	}
}