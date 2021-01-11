package test_fib_bench_demo

import "testing"

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

/*
测试所有：
D:\goproject\src\base_stu\test_fib_demo>go test -bench=.
goos: windows
goarch: amd64
pkg: base_stu/test_fib_demo
BenchmarkFib1-8         672065593                1.78 ns/op
BenchmarkFib2-8         237319522                5.08 ns/op
BenchmarkFib3-8         137665387                8.68 ns/op
BenchmarkFib10-8         3831884               315 ns/op
BenchmarkFib20-8           30692             39091 ns/op
BenchmarkFib40-8               2         590439100 ns/op
PASS
ok      base_stu/test_fib_demo  10.094s



测试其中一个：
D:\goproject\src\base_stu\test_fib_demo>go test -bench=Fib3
goos: windows
goarch: amd64
pkg: base_stu/test_fib_demo
BenchmarkFib3-8         138050074                8.63 ns/op
PASS
ok      base_stu/test_fib_demo  2.098s


*/

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }