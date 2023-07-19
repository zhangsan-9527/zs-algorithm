package main

import "fmt"

func f(s []int, str string) {

	str = "你好"
	// 对传进来的数组进行修改
	s[0] = 8888
	s[2] = 99999
	//fmt.Printf("s地址:%p\n", &s) // s地址:0xc000008090
	fmt.Println(str) // 你好

}

func main() {
	str := "gogogo"
	sli := make([]int, 3)
	f(sli, str)
	//fmt.Printf("s地址:%p", &sli) // s地址:0xc000008078   证明s, sli这是两个地址  但是每个元素的指针是相通的, 里面的每个元素是指针
	fmt.Println(str) // gogogo
}
