package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
	存储全局参数，供其他模块使用
	通过conf.json配置
*/

type GlobalObj struct {
	Host  string
	Port  int
	Nagle bool
}

var GlobalObject *GlobalObj

func init() {

	//默认值
	GlobalObject = &GlobalObj{
		Host:  "0.0.0.0",
		Port:  8888,
		Nagle: true,
	}
	GlobalObject.Reload()
}

func (G *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("./conf/conf.json") //window环境下改为"./conf/conf.json"
	if err != nil {
		fmt.Println("Read file err,", err)
		panic(err.(interface{}))
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("Read file err,", err)
		panic(err.(interface{}))
	}
}
