package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct {
	Name string	`json:"name"`
}

type Client struct{
	send chan Message
}

func (client *Client) write(){
	// range keword used for pull single value from client.send
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
	}
}

/**
 * Generates a random number between 0 and 10
 */
func r() time.Duration{
	return time.Second * time.Duration(rand.Intn(10))
}

func (client *Client) subscribeChannels(){
	for {
		time.Sleep(r())
		client.send <- Message{"channel add"}
	}
}

func (client *Client) subscribeMessages(){
	for {
		time.Sleep(r())
		client.send <- Message{"message add"}
	}
}

func NewClient() *Client{
	return &Client{
		send: make(chan Message),
	}
}

func main(){
	client := NewClient()
	go client.subscribeChannels()
	go client.subscribeMessages()
	client.write()
}