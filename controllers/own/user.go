package own

import (
	"fmt"
	"github.com/astaxie/beego"
	"myprojectapi/models/common"
	"myprojectapi/models/own"
)

// Operations about Users
type UserController struct {
	beego.Controller
}


// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if username == "" || password == ""  {
		u.Data["json"] = common.Fail(200,"用户名和密码不能为空")
		goto endPoint
	}
	fmt.Println(own.Login(username, password))
	if fuser,err:=own.Login(username, password);err!=nil {
		u.Data["json"] = common.Fail(200,"用户名或密码错误")
	} else {
		token := own.GenerateToken()
		u.SetSession(token, fuser.FLoginName)
		user := map[string]interface{}{"token":token,"userinfo":fuser}
		u.Data["json"] = common.Success(user)
	}

	endPoint:
		u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	token := u.GetString("token")
	u.DelSession(token)
	u.Data["json"] = common.SuccessMsg("退出登陆成功")
	u.ServeJSON()
}

func (u *UserController) Getlist(){
	userName := u.GetString("username")
	page,err := u.GetInt("page",1)
	if err != nil{
		panic(err)
	}
	pageSize,err := u.GetInt("pageSize",10)
	if err != nil{
		panic(err)
	}
	params := map[string]interface{}{"fnickname":userName}

	userList := own.GetUserList(page, pageSize, params)
	u.Data["json"] = common.Success(userList)
	u.ServeJSON()
}

func (u *UserController) List(){
	u.Data["json"] = common.SuccessMsg("测试")
	u.ServeJSON()
}

