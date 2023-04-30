package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string	`json:"name"`
	Data interface{}	`json:"data"`
}

type Speaker interface {
	Speak()
}

type Channel struct {
	Id string	`json:"id"`
	Name string	`json:"name"`
}

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

func addChannel(data interface{}) (Channel, error) {
	fmt.Println("Add Channel Function")
	fmt.Println(data)

	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil{
		return channel, err
	}

	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return channel, nil
}