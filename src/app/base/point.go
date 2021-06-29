package util

import "fmt"

/**

普通变量： 存储数据本身
指针变量： 存储值的内存地址

*/
func TestPoint() {
	// m1()
	// m2()
	m3()

}

func m1() {
	aint := 1
	ptr := &aint
	println(ptr)
	println(*ptr)
}

func m2() {
	astr := new(string)
	*astr = "测试指针"

	println(astr)
	println(*astr)
}

func m3() {
	aint := 1
	var bint *int
	bint = &aint

	fmt.Printf("%p \n", bint)
	println(bint)
	println(*bint)

	println(aint)
	println(&aint)
}
