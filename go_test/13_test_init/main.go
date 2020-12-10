package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 编写代码利用反射实现一个ini文件的解析器程序。

// mysql配置结构的结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Test     bool   `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0.参数的校验
	// 0.1 传进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") // 新建一个错误
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针（因为配置文件中各种键值对对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct") // 新建一个错误
		return
	}
	// 1.读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b)   // 将文件内容转成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)

	// 2.一行一行得读数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串收尾的空格
		line = strings.TrimSpace(line)
		// 2.1 如果是注释就跳过,或空行也就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		// 2.2 如果是[开头的就表示节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 把这一行收尾去掉，取到中间内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			//v := reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体,把字段名记录下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s", sectionName, structName)
				}
			}

		} else {
			// 2.3 如果不是[开头就是=分割的键值对
			// （1）以等号分割这一行，等号左边是key,右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d systax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index]) // 结构体的key
			kValue := strings.TrimSpace(line[index+1:])
			// （2）根据structName去data里面对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fileName string // 配置文件=号前面的值
			var fileType reflect.StructField
			// （3）遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sType.NumField(); i++ {
				filed := sType.Field(i) // tag 信息是存储在
				fileType = filed
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					fileName = filed.Name
					break // 如果没有break，出现重名的话，变量会用后面的
				}
			}
			// 4.如果key = tag 给这个字段赋值
			// (1) 根据keyName 去取出这个字段
			if len(fileName) == 0 {
				// 在结构体中找不到对应的字符
				continue
			}
			fileObj := sValue.FieldByName(fileName)

			// （2）对其赋值
			fmt.Println(fileName, fileType.Type.Kind()) // ,valueName.Type().Kind()
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(kValue)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(kValue, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(kValue, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(kValue)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)

			}
		}

	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}
	fmt.Printf("%#v", cfg)
	// 第三方ini解析器
	// https://github.com/go-ini/ini
}
