package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myprojectapi/models/common"
	_ "myprojectapi/routers"
	"strings"
	"time"
)



func init(){
	common.GetEnv()
	mysql_user := beego.AppConfig.String("mysql_user")
	mysql_pass := beego.AppConfig.String("mysql_pass")
	mysql_db := beego.AppConfig.String("mysql_db")
	mysql_host := beego.AppConfig.String("mysql_host")
	mysql_port := beego.AppConfig.String("mysql_port")


	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{mysql_user, ":", mysql_pass, "@tcp(",mysql_host, ":", mysql_port, ")/", mysql_db, "?charset=utf8"}, "")
	orm.RegisterDataBase("default","mysql",path)
}

func main() {
	fmt.Println(beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.BConfig.WebConfig.Session.SessionOn = true
	//获取当前的日期（月、日）
	fileName := time.Now().Format("0102")
	
	common.GetCurrentDirectory()
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogger("file", `{"filename":"logs/info`+fileName+`.log"}`)
	beego.SetLogFuncCall(true)
	beego.Notice("this is a notice information")
	beego.Run()
}



