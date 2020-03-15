package main

import "fmt"
import "strconv"
import "strings"

func main() {

	// 统计字符串的长度，按字节
	// 一个英文字母占用一个字节
	// utf8中一个中文编码一个字符三个字节
	str := "hello东"
	fmt.Println("str len=", len(str)) // 8

	str2 := "hello深圳"
	// 有中文的话，需要转rune
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符=%c\n", str3[i])
	}

	// 字符串转整数：n,err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	// 整数转字符串
	str4 := strconv.Itoa(12345)
	fmt.Printf("str=%v, str_type=%T\n", str4, str4)

	// 字符串转byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n",bytes)

	// byte 转字符串
	str5 := string([]byte{97, 98, 99})
	fmt.Println("str5=", str5)

	// 10 进制转 2,8,16进制  返回对应的字符串
	str6 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n",str6)
	str7 := strconv.FormatInt(123,8)
	fmt.Printf("123对应的8进制是=%v\n",str7)
	str8 := strconv.FormatInt(123,16)
	fmt.Printf("123对应的16进制是=%v\n",str8)

	// 查找子串是否在指定的字符串中
	is_in := strings.Contains("hszllo","llo")
	fmt.Println("是否存在",is_in)

	// 统计一个字符串中有几个指定的子串
	num := strings.Count("hahahahahhahah", "ha")
	fmt.Printf("出现%v次\n",num)

	// 不区分带小写比较
	is_eq := strings.EqualFold("ASD","asd")
	fmt.Println("是否相同",is_eq)

	// 返回子串Index值
	index := strings.Index("NNN_ABC","ABC")
	fmt.Println("返回index=",index)

	// 返回子串最后一次出现的index,如果没有返回-1
	index_02 := strings.LastIndex("go go go goLang","go")
	fmt.Println("最后一次出现index=",index_02)

	//字符串的替换
	rp_str := strings.Replace("hello go go go","go","go 语言",3) 
	fmt.Println("替换后的结果:",rp_str)

	// 字符串的分割，才分为字符串数组
	strArr := strings.Split("hello,world,ok",",")
	fmt.Println("strArr=",strArr)
	// 遍历数组
	for i := 0 ; i <len(strArr); i++ {
		fmt.Println(strArr[i])
	}

	// 字符串中的带小写转换
	// strings.ToLower("GO")
	// strings.ToUpper("go")
	res1 := strings.ToLower("GO")
	fmt.Println("变成小写：",res1)
	res2 := strings.ToUpper("go")
	fmt.Println("变成大写：",res2)

	// 字符串左右空格去掉
	new_str_01 := strings.TrimSpace("  hello go !   ")
	fmt.Println("去掉空格的:",new_str_01)

	// 去掉左右两边指定的字符串
	new_str_02 := strings.Trim("!!hello go !!","!!")
	fmt.Println("去掉左右两边指定的！:",new_str_02)

	// 去掉左边指定的字符串
	new_str_03 := strings.TrimLeft("!!hello go !   ","!!")
	fmt.Println("去掉左边指定的！:",new_str_03)

	// 去掉右边指定的字符串
	new_str_04 := strings.TrimRight("!!hello go !!","!!")
	fmt.Println("去掉右边指定的！:",new_str_04)

	// 判断字符串是否以指定的字符串开头
	is_start := strings.HasPrefix("ftp://127.0.0.1","ftp")
	fmt.Println("是否为ftp开头:",is_start)

	// 判断字符串是否以指定的字符串结尾
	is_end := strings.HasSuffix("www.baidu.com","com")
	fmt.Println("是否为com结尾:",is_end)


}
