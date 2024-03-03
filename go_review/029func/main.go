package main

import (
	"fmt"
	"strings"
)

func countWords(s string) map[string]int {
	// 将字符串按空格分割成单词切片
	//words := strings.Fields(s)
	words := strings.Split(s, " ")
	fmt.Println(words)
	// 创建一个用于存储单词计数的map
	wordCount := make(map[string]int)

	// 遍历单词切片，统计每个单词出现的次数
	for _, word := range words {
		// 将每个单词转换为小写，以便不区分大小写地统计
		word = strings.ToLower(word)

		// 如果单词已经在map中，则增加计数
		if count, ok := wordCount[word]; ok {
			wordCount[word] = count + 1
		} else {
			// 如果单词不在map中，则添加到map并设置计数为1
			wordCount[word] = 1
		}
	}

	return wordCount
}

func main() {
	input := "how do you do"
	wordCounts := countWords(input)

	// 打印每个单词及其出现的次数
	for word, count := range wordCounts {
		fmt.Printf("%s=%d\n", word, count)
	}
}
