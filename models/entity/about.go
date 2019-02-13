package entity

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
)

var (
	FAboutTypes map[string]*FAboutType
)
type FAboutType struct{
	Fid int64 `orm:"column(fid);pk" json:"fid"`
	Fsort int64 `orm:"column(fsort)" json:"fsort"`
	Fstate byte `orm:"column(fstate)" json:"fstate"`
	Ftitle string `orm:"column(ftitle)" json:"ftitle"`
	Fdescribe string `orm:"column(fdescribe)" json:"fdescribe"`
	Fagentid int64 `orm:"column(fagentid)" json:"fagentid"`
}

type FAgent struct {
	Fid int64 `orm:"column(fid);pk" json:"fid"`
	Fname string `orm:"column(fname)" json:"fname"`
	Fphone string `orm:"column(fphone)" json:"fphone"`
	Fdomain string `orm:"column(fdomain)" json:"fdomain"`
	Fremark string `orm:"column(fremark)" json:"fremark"`
	//Fcreatetime time.Time `orm:"column(fcreatetime)" json:"fcreatetime"`
	Fcreatetime string `orm:"column(fcreatetime)" json:"fcreatetime"`
}



func SelectFAgentOne(id int) FAgent{
	o := orm.NewOrm()
	o.Using("default")
	var fagent FAgent
	o.QueryTable("FAgent").Filter("fid",id).One(&fagent)
	UpdateTime("2018-09-09 12:38:50")
	return  fagent
}

func UpdateTime(date string) FAgent{
	o := orm.NewOrm()
	o.Using("default")
	var fagent FAgent
	beego.Notice("this is notice")
	o.Raw("update f_agent set fcreatetime = ? where fid = ?" ,date, 1).Exec()
	return  fagent
}






