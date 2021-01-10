package test_split_string

import (
	"reflect"
	"testing"
)

// 在test文件下运行：go test
// 会执行这个目录下所有的测试用例
func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接受一个*testing.T 类型参数
	got := Split("a:b:c", ":")      // 程序输出的结构
	want := []string{"a", "b", "c"} // 期望的结果
	if !reflect.DeepEqual(want, got) {
		// 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v,got:%v", want, got) // 测试失败输出错误提示
	}
}

func Test2Split(t *testing.T) {
	got := Split("abcdf", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v,got:%v", want, got)
	}
}
