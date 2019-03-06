package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	// "net/http"
	// "strconv"
	"time"
)


type Client struct{
	id        string
	socket    *websocket.Conn
	send      chan []byte
	userId    string
}


type Message struct{
	  Sender   	string `json:"sender,omitempty"`
	  Recipient string `json:"recipient,omitempty"`
	  Content   string `json:"content,omittempty"`
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	// ptp        chan mess
}
var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
	// ptp:    	  make(chan mess),
}

func init(){
	go manager.start()
	go manager.broast()
}

func (manager *ClientManager) broast(){
	for {
		time.Sleep(10*time.Second)
		manager.broadcast <- []byte("我是广播消息")
		
	}
}


func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		// case message := <-manager.ptp:
		// 	for conn := range manager.clients{
		// 		if strconv.Itoa(message.SendId) == conn.userId || strconv.Itoa(message.RecvId) == conn.userId {
		// 			jsonMessage,_ :=json.Marshal(message)
		// 			conn.send <- jsonMessage
		// 		}
		// 	}

		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore  {
			conn.send <- message
		}
	}
}





type mess struct {
	Message string
	MessageType string
	SendId  int
	RecvId  int
	Time    string
}

func (c *Client) Read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		var messes mess
		json.Unmarshal(message,&messes)
		fmt.Println("broadcast")
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
		// if messes.RecvId == 0 {
		// 	fmt.Println("broadcast")
		// 	jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		// 	manager.broadcast <- jsonMessage
		// }else {
		// 	fmt.Println("other")
		// 	manager.ptp <- messes
		// }

	}
}

func (c *Client) Write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			fmt.Println("写入信息：",string(message))
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func WsHandler(id,userId string,conn *websocket.Conn){

	client := &Client{id: id,userId:userId , socket: conn, send: make(chan []byte)}
	manager.register <- client
	
	go client.Read()
	go client.Write()
}
