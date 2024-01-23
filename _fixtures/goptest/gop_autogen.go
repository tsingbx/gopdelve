package main

import (
	"fmt"
	"os"
	"gop-sample/cpkag"
	ab "gop-sample/cpkag/b"
	"gop-sample/cpkag/b/ss"
	"gop-sample/cpkag/k"
	"gop-sample/cpkag/utils"
	"gop-sample/gpkag"
	utils1 "gop-sample/utils"
	"path/filepath"
	utils2 "github.com/liuscraft/testgomod/utils"
)

const _ = true
//line a.gop:3:1
func a() {
//line a.gop:4:1
	fmt.Println("a")
}
//line a.gop:7:1
func Ab() {
//line a.gop:8:1
	fmt.Println("ab")
}
//line b.gop:18
func main() {
//line b.gop:18:1
	bb := 3
//line b.gop:19:1
	Ab()
//line b.gop:20:1
	ab.Ab()
//line b.gop:21:1
	ab.Ac()
//line b.gop:22:1
	ss.Ab()
//line b.gop:23:1
	ss.Bs()
//line b.gop:24:1
	fmt.Println(bb)
//line b.gop:25:1
	k.K()
//line b.gop:26:1
	k.Kk()
//line b.gop:27:1
	k.Ab()
//line b.gop:28:1
	utils2.TestCsgo()
//line b.gop:29:1
	utils.TestCsgo()
//line b.gop:30:1
	utils1.TestCsgo()
//line b.gop:32:1
	a()
//line b.gop:33:1
	cpkag.P()
//line b.gop:34:1
	cpkag.Gg()
//line b.gop:35:1
	gpkag.G()
//line b.gop:36:1
	fmt.Println("b")
//line b.gop:37:1
	fmt.Println("b")
//line b.gop:38:1
	fmt.Println("b")
//line b.gop:39:1
	fmt.Println("b")
//line b.gop:40:1
	fmt.Println("b")
//line b.gop:41:1
	fmt.Println("b")
//line b.gop:42:1
	fmt.Println("b")
//line b.gop:45:1
	absolutePath := "/Users/liushunshun/codings/goplus_project/gop-sample"
//line b.gop:48:1
	relativePath := "../"
//line b.gop:51:1
	fullPath := filepath.Join(absolutePath, relativePath)
//line b.gop:54:1
	fileInfo, err := os.Stat(fullPath)
//line b.gop:55:1
	if err != nil {
//line b.gop:56:1
		if os.IsNotExist(err) {
//line b.gop:57:1
			fmt.Println("文件不存在")
		} else {
//line b.gop:59:1
			fmt.Println("发生错误:", err)
		}
//line b.gop:61:1
		return
	}
//line b.gop:64:1
	if fileInfo.IsDir() {
//line b.gop:65:1
		fmt.Println("找到目录:", fullPath)
	} else {
//line b.gop:67:1
		fmt.Println("找到文件:", fullPath)
	}
}
