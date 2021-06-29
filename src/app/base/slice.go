package util

import "fmt"

func TestSlice() {

	s1()
}

func s1() {
	myarr := [...]int{1, 2, 3, 4, 5}

	fmt.Println(myarr)

	myslice1 := myarr[1:3]
	myslice2 := myarr[1:3:4]
	fmt.Println(myslice1)
	fmt.Println(myslice2)
}

func s2() {
	var strlist []string
	var numlist []int
	var numListEmpty = []int{}
}
