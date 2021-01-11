package test_group_demo

import (
	"reflect"
	"testing"
)

// 测试组的应用

/*

go test -run="TestGroupSplit/case_3"  跑测试组里面的单个测试
D:\goproject\src\base_stu\split_string>go test -run="TestGroupSplit/case_3"
PASS
ok      base_stu/split_string   0.026s

Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件
go test -cover -coverprofile=c.out
*/
func TestGroupSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接受一个*testing.T 类型参数
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case_1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": {"dqudqb", "u", []string{"dq", "dqb"}},
		"case_4":{"南山有山有什么", "山", []string{"南","有","有什么"}},
	}
	// 遍历切片，逐一执行测试用例
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}

}

// 基准测试
func BenchmarkSplit(b *testing.B){
	for i :=0;i<b.N;i++{
		Split("a.b.c.d",".")
	}
}