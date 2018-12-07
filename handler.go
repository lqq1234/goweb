package handler

import (
	"io/ioutil"
	"log"
	"os"
	"io"
)
// 定义日志
// 定义四种日志
var(
	Trace *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)
// 初始化
// 初始化日志
func init(){
	file,err := os.OpenFile("../error.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil{
		log.Fatal("Failed to open error log file:", err)
	}
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout,
		"Warning: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr),
		"Error: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
//asdfjkdsjk
//asdfdsafdsa
//sadfdsafsafss
//sdafdsafs
// 获取配置文件信息
//返回值是byte数组
func GetConfigContent()[]byte{
//asdjfkjsak
	conf,_:=ioutil.ReadFile("/opt/conf/goconf")
	if len(conf) == 0 {
		conf = []byte("This is default value of goconf.")
	}
//sadfkjsak
	return conf
}
///sadfkjdasj
//sadfdsakf
//adsfdsaf
//用于测试的方法
func Add(a,b int)int{
	return a+b
}
