package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gotest/myprojectapi/models/entity"
	"strconv"
)

var (
	Abouts map[string]*About
	FAbouts map[string]*FAbout
)

type About struct {
	Fid   int64		   `orm:"column(fid);pk" json:"fid"`
	Fabouttype  int64  `orm:"column(fabouttype)" json:"fabouttype"`
	Fshortname string  `orm:"column(fshortname)" json:"fshortname"`
	Fsort string		`orm:"column(fsort)" json:"fsort"`
	Ftitle string		`orm:"column(ftitle)" json:"ftitle"`
	Fcontent string		`orm:"column(fcontent)" json:"fcontent"`
	Fshowid string		`orm:"column(fshowid)" json:"fshowid"`
}

type FAbout struct {
	Fid   int64		   `orm:"column(fid);pk" json:"fid"`
	Fabouttype  int64  `orm:"column(fabouttype)" json:"fabouttype"`
	Fshortname string  `orm:"column(fshortname)" json:"fshortname"`
	Fsort string		`orm:"column(fsort)" json:"fsort"`
	Ftitle string		`orm:"column(ftitle)" json:"ftitle"`
	Fcontent string		`orm:"column(fcontent)" json:"fcontent"`
	Fshowid string		`orm:"column(fshowid)" json:"fshowid"`
}


func SelectAboutOne(id string) (about About) {
	o := orm.NewOrm()
	o.Using("default")
	strtoint,err := strconv.ParseInt(id,10,64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("this is a about object %v" ,about)
	about.Fid = strtoint
	fabout := new(FAbout)
	//aa :=o.Read(&about)
	fmt.Printf("this is a about object111 %v \n" ,fabout)
	fmt.Println(fabout)

	qs := o.QueryTable("f_about").Filter("fid",id).One(fabout)
	fmt.Printf("this is a about object312312 %v \n" ,fabout)
	fmt.Printf("this is a about object312312 %v \n" ,qs)
	r := o.Raw("select * from f_about where fid = ?", id)
	errs := r.QueryRow(&about)
	if errs != nil {
		panic(errs)
	}
	fmt.Printf("this is a about object %v \n" ,about)
	fmt.Printf("this is a about object %v \n" ,r)

	//fmt.Printf("this is a about object %v" ,aa)
	return about
}
func init(){
	orm.RegisterModel(new(About),new(FAbout),new(entity.FAgent))
}




