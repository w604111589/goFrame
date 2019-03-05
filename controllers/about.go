package controllers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"goFrame/models"
	"goFrame/models/common"
	"goFrame/models/entity"

	"github.com/astaxie/beego"
)

// Operations about object
type AboutController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *AboutController) GetAboutOne() {
	var ob models.About
	//json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	id := o.GetString("id")
	idInt,_ := o.GetInt("id")
	idInput := o.Input().Get("id")
	fmt.Println(ob)
	fmt.Println(id)
	fmt.Println(idInt)
	fmt.Println(idInput)
	ob = models.SelectAboutOne(id)
	o.Data["json"] =ob
	o.ServeJSON()
}

func (o *AboutController) GetFAgentOne() {
	id,_ := o.GetInt("id")
	beego.Notice("this is a fagentone information")

	rc := common.RedisClient.Get()
	defer rc.Close()

	//fmt.Printf("31323: %v",rc)
	fmt.Printf("31323: %v \n",common.RedisClient)
	rc.Do("SET","name","wangtao32131")
	res,_ := redis.String(rc.Do("GET","name1"))
	res1,_ := redis.String(rc.Do("GET","name"))

	fmt.Println("31323:",res)
	fmt.Println("31323:",res1)
	fmt.Println(res)
	fmt.Println(id)
	ob := entity.SelectFAgentOne(id)
	o.Data["json"] =ob
	o.ServeJSON()
}

/**
 *
 */

