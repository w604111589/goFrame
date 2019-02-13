package own

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
	"strconv"
	"time"
)


type FUser struct{
	Fid				int 	`orm:"column(fid);pk" json:"fid" `
	FLoginName		string 	`orm:"column(floginname);pk" json:"floginname" `
	FNickname		string 	`orm:"column(fnickname)" json:"fnickname" `
	FLoginPassword	string 	`orm:"column(floginpassword)" json:"floginpassword" `
	FTelephone		string 	`orm:"column(ftelephone)" json:"ftelephone" `
	FEmail			string 	`orm:"column(ftelephone)" json:"femail" `
}

func getUserOne(username  string) *FUser{
	o := orm.NewOrm()
	o.Using("default")
	var fuser *FUser
	o.Raw("select * from f_user where floginname=? ",username).QueryRow(&fuser)
	return fuser
}

func GetUserList(page,pageSize int,filters map[string]interface{}) (FUsers []*FUser){
	offset :=  (page - 1)*pageSize
	query := orm.NewOrm().QueryTable("f_user")
	if len(filters)> 0{
		for key,value := range filters{
			query = query.Filter(key,value)
		}
	}
	query.Count()
	query.OrderBy("-fid").Limit(pageSize,offset).All(&FUsers)
	return FUsers
}

func Login(username, password string) (fuser *FUser,err error) {
	fuser =  getUserOne(username)
	if fuser.FLoginPassword== Md5(password){
		return fuser,nil
	}
	return nil,errors.New("用户名或密码有误")
}



func Md5(str string) string {
	crypto := md5.New()
	crypto.Write([]byte(str))
	return hex.EncodeToString(crypto.Sum(nil))
}

func GenerateToken() string{
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(time.Now().Unix(), 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
