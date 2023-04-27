package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string
	Data interface{}
}

type Speaker interface {
	Speak()
}

type Channel struct {
	Id string
	Name string
}

func main(){
	recRawMsg := []byte (`{"name":"channel add", "data":{"name":"Hardware Support"}}`)

	var recMessage Message
	err := json.Unmarshal(recRawMsg, &recMessage)
	
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", recMessage)	// %#v prints out the struct with the field names

	if recMessage.Name == "channel add"{
		addChannel(recMessage.Data)
	}


}

func addChannel(data interface{}) (Channel, error) {
	fmt.Println("Add Channel Function")
	fmt.Println(data)

	var channel Channel
	channelMap := data.(map[string]interface{})
	channel.Name = channelMap["name"].(string)
	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return channel, nil
}