
package controllers

import (
	"log"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"goFrame/models"
)

type ChatController struct{
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func  (c *ChatController) Ws()  {

	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
    	if err != nil {
        log.Fatal(err)
    	}

	userId := c.getString("userId","0")

	if error != nil {
		http.NotFound(res, req)
		return
	}
	newV4,_ := uuid.NewV4()
	fmt.Println("V4:",newV4.String())
	client := &Client{id: newV4.String(),userId:userId , socket: conn, send: make(chan []byte)}

	manager.register <- client

	go client.read()
	go client.write()
	
}