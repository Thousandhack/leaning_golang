package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// 编写代码利用反射实现一个ini文件的解析器程序。

// mysql配置结构的结构体
type MysqlConfig struct {
	Address string `ini:"address"`
	Port int	`ini:"port"`
	Username string  `ini:"username"`
	Password string  `ini:"password"`
}


type RedisConfig struct {
	Host string `ini:"host"`
	Port int	`ini:"port"`
	Password string  `ini:"password"`
	Database string `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string,data interface{})(err error){
	// 0.参数的校验
	// 0.1 传进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t,t.Kind())
	if t.Kind() != reflect.Ptr{
		err = errors.New("data should be a pointer")  // 新建一个错误
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针（因为配置文件中各种键值对对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct{
		err = errors.New("data should be a struct")  // 新建一个错误
		return
	}
	// 1.读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil{
		return
	}
	//string(b)   // 将文件内容转成字符串
	lineSlice := strings.Split(string(b),"\r\n")
	fmt.Printf("%#v\n",lineSlice)

	// 2.一行一行得读数据
	for idx,line := range lineSlice{
		// 去掉字符串收尾的空格
		line = strings.TrimSpace(line)
		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line,";") || strings.HasPrefix(line,"#"){
			continue
		}
		// 2.2 如果是[开头的就表示节（section）
		if strings.HasPrefix(line,"["){
			if line[0] != '[' || line[len(line)-1] != ']'{
				err = fmt.Errorf("line:%d syntax error",idx+1)
				return
			}
			// 把这一行收尾去掉，取到中间内容
			sectionName := strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName) == 0{
				err = fmt.Errorf("line:%d syntax error",idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			// 91 15分钟

		} else {
			// 2.3 如果不是[开头就是=分割的键值对
		}

	}

	return
}

func main()  {
	var cfg Config
	err := loadIni("./conf.ini",&cfg)
	if err != nil{
		fmt.Printf("load ini failed,err:%v\n",err)
		return
	}
	//fmt.Println(cfg.Address,cfg.Port,cfg.Username,cfg.Password)
}
