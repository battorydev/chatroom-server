package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/mitchellh/mapstructure"
	"time"
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
		var inMessage Message
		var outMessage Message
		err := socket.ReadJSON(&inMessage)
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)

		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil{
				outMessage = Message{"error", err}	//????
				err := socket.WriteJSON(outMessage)
				if err != nil{
					fmt.Println(err)
					break
				}
			}
			//TODO call database function

			// if you want to format the output, you can use the keyword 'fallthrough' (uncomment the line below)
			// fallthrough
		case "channel subscribe":
			subscribeChannel(socket)
			//TODO call database function
		}
	}
}

func addChannel(data interface{}) (error) {
	fmt.Println("Add Channel Function")
	fmt.Println(data)

	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil{
		return err
	}

	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	fmt.Println("Added Channel.")
	return nil
}

func subscribeChannel(socket *websocket.Conn){

	//TODO call database function (changefeed)
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add", Channel{"1", "Software Support"}}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}