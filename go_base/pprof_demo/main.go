package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:  // 阻塞
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool // 是否开启CPUprofile的标志位
	var isMemPprof bool // 是否开启内存profile的标志位

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

/*
	使用命令：go run main.go -cpu=true -mem=true   生成pprof文件
	然后再用pprof工具来进行分析文件

	D:\goproject\src\base_stu\pprof_demo>go tool pprof cpu.pprof
	Type: cpu
	Time: Jan 11, 2021 at 2:04pm (CST)
	Duration: 20.20s, Total samples = 2.67mins (792.19%)
	Entering interactive mode (type "help" for commands, "o" for options)
	(pprof) top 3
	Showing nodes accounting for 159.02s, 99.39% of 159.99s total
	Dropped 56 nodes (cum <= 0.80s)
		  flat  flat%   sum%        cum   cum%
		76.69s 47.93% 47.93%    129.71s 81.07%  runtime.selectnbrecv
		52.97s 33.11% 81.04%     52.98s 33.11%  runtime.chanrecv
		29.36s 18.35% 99.39%    159.07s 99.42%  main.logicCode
	(pprof) top 5
	Showing nodes accounting for 159.02s, 99.39% of 159.99s total
	Dropped 56 nodes (cum <= 0.80s)
		  flat  flat%   sum%        cum   cum%
		76.69s 47.93% 47.93%    129.71s 81.07%  runtime.selectnbrecv
		52.97s 33.11% 81.04%     52.98s 33.11%  runtime.chanrecv
		29.36s 18.35% 99.39%    159.07s 99.42%  main.logicCode
	(pprof) quit

	D:\goproject\src\base_stu\pprof_demo>go run main.go -cpu=true -mem=true

	D:\goproject\src\base_stu\pprof_demo>go tool pprof cpu.pprof
	Type: cpu
	Time: Jan 11, 2021 at 2:12pm (CST)
	Duration: 20s, Total samples = 30ms ( 0.15%)
	Entering interactive mode (type "help" for commands, "o" for options)
	(pprof) top 4
	Showing nodes accounting for 30ms, 100% of 30ms total
	Showing top 4 nodes out of 16
		  flat  flat%   sum%        cum   cum%
		  10ms 33.33% 33.33%       10ms 33.33%  runtime.(*timersBucket).addtimerLocked
		  10ms 33.33% 66.67%       10ms 33.33%  runtime.procyield
		  10ms 33.33%   100%       10ms 33.33%  runtime.stdcall1
			 0     0%   100%       10ms 33.33%  main.logicCode
	(pprof) top 5


*/
