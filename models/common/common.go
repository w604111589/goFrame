package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)


/**
 * @author    wangtao
 * @date      2018/11/1 17:53
 * @param     nil
 * @return    string
 * @desc      根据日期每天创建日志文件
 */
func GetCurrentDirectory() string{
	dir , err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Notice(err)
	}
	CreatePath(dir+"/logs")
	// 获取当前日期的月日
	fileName := time.Now().Format("0102")
	CreateFile(dir+"/logs/info"+fileName+".log")
	return dir
}

 /*
  * @author wangtao
  * @date   2018/11/1 14:26
  * @param  string
  * @return (bool,error)
  */
func PathExists(path string) (bool,error){

	_,err := os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return true,err
}

/**
 * @author    wangtao
 * @date      2018/11/1 17:51
 * @param     string filePath:文件的完整路径
 * @return    bool
 * @desc      文件是否存在，路径不存在创建路径，文件不存在，则创建文件
 */
func CreateFile(filePath string) (bool){
	res,_ := PathExists(filePath)
	if !res{
		_,err := os.Create(filePath)
		fmt.Println(err)
		if err != nil{
			fmt.Println(err)
			return false
		}else{
			fmt.Println("创建文件成功")
		}

	}

	return true
}


/**
 * @author    wangtao
 * @date      2018/11/1 17:49
 * @param     string path:本机路径
 * @return    bool
 * @desc      路经是否存在在不存在则创建
 */
func CreatePath(path string) (bool){
	res,_ := PathExists(path)
	if !res{
		rs := os.MkdirAll(path,os.ModePerm)
		if rs != nil{
			fmt.Println(rs)
			return false
		}
	}
	return true
}


/**
 * @author    wangtao
 * @date      2018/11/1 17:49
 * @param     nil
 * @return    nil
 * @desc      调用cmd命令
 */
func GetIpAddress(){
	www , _ := exec.Command("CMD","/C","ipconfig").Output()
	fmt.Println(string(www))
}

/**
 * @author    wangtao
 * @date      2018/11/1 17:49
 * @param     nil
 * @return    string
 * @desc      获取本机的内网IP
 */
func GetLocalIp() string {
	addrSlice, _:= net.InterfaceAddrs()
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr := ipnet.IP.String()
				fmt.Println("printIp",IpAddr)
				return IpAddr
			}
		}
	}
	IpAddr := "localhost"
	fmt.Println("printIp",IpAddr)
	return IpAddr
}




 /**
  * @author wangtao
  * @date   2018/11/1 17:37
  * @param  nil
  * @return  string
  */
func GetEnv() string{
	localIp := GetLocalIp()
	switch localIp{
	case  "172.31.22.71":
		beego.BConfig.RunMode = "test"
	default:
		beego.BConfig.RunMode = "dev"
	}
	return beego.BConfig.RunMode
}
