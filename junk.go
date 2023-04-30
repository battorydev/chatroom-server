package main

import (
	"fmt"
	"encoding/json"
)





func main1(){
	recRawMsg := []byte (`{"name":"channel add", "data":{"name":"Hardware Support"}}`)

	var recMessage Message
	err := json.Unmarshal(recRawMsg, &recMessage)
	
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", recMessage)	// %#v prints out the struct with the field names

	if recMessage.Name == "channel add"{
		channel, err := addChannel(recMessage.Data)
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMsg, err := json.Marshal(sendMessage)	// returns a byte array
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(string(sendRawMsg))
	}


}

