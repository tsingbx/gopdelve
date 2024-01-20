package main

import (
	"fmt"
	"testgop/utils"
	"github.com/liuscraft/testgomod"
	utils1 "github.com/liuscraft/testgomod/utils"
)
//line main.gop:9:1
func test() {
//line main.gop:10:1
	a := 1
//line main.gop:11:1
	defer func() {
//line main.gop:12:1
		fmt.Println(a)
	}()
//line main.gop:15:1
	utils1.TestCsgo()
//line main.gop:16:1
	utils.TestCsgo()
//line main.gop:17:1
	a = 2
}
//line main.gop:20
func main() {
//line main.gop:20:1
	test()
//line main.gop:21:1
	Version()
//line main.gop:22:1
	fmt.Println(testgomod.Version())
}
