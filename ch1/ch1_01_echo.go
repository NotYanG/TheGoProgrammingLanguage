// 输出其命令行参数

package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	// echo1输出其命令行参数
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	// echo2输出其命令行参数
	// 另一种形式的for循环
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	/* 使用strings包中的Join函数 */
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {

}
