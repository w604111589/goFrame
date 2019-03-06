
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
	"goFrame/models/common"
	"fmt"
	"goFrame/models"
)

type ChatController struct{
	beego.Controller
}

var upgrader = websocket.Upgrader{CheckOrigin:func(r *http.Request) bool{return true}}


func (c *ChatController) Test(){
	fmt.Println("22222222222222222")
	c.Data["json"] = common.SuccessMsg("请求成功")
	c.ServeJSON()
}

func  (c *ChatController) Ws()  {
	userId := c.GetString("userId","0")

	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
    	if err != nil {
        beego.Error(err)
    	}
	fmt.Println("连接已建立，用户ID为：",userId)
	newV4,_ := uuid.NewV4()

	models.WsHandler(newV4.String(),userId,conn)

	
}