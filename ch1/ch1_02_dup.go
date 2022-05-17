// 输出标准输入中出现次数大于1的行，前面是次数

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup1() {
	// dup1输出标准输入中出现次数大于1的行，前面是次数
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++

		/* counts[input.Text()]++ 等价以下代码 */
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}

	// 注意：忽略input.Err()中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	// dup2打印输入中多次出现的行的个数和文本
	// 它从stdin或指定的文件列表读取
	counts := make(map[string]int)
	args := os.Args[1:]
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			file, err := os.Open(arg) // 打开文件
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			err = file.Close()
			if err != nil {
				return
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意：忽略input.Err()中可能的错误
}

func dup3() {
	// 使用ReadFile读取文件
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	//dup1()
	//dup2()
	dup3()
}
