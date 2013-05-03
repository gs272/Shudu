package main

import (
  . "box"
	"fmt"
	file "io/ioutil"
)

var formlist = Box_Num{
	{9, 0, 0, 0, 2, 0, 4, 0, 1},
	{0, 0, 5, 0, 9, 0, 0, 2, 0},
	{0, 4, 0, 1, 0, 0, 0, 9, 7},
	{5, 0, 0, 0, 0, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 2},
	{4, 0, 0, 5, 7, 0, 3, 0, 9},
	{0, 5, 3, 0, 0, 2, 1, 8, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 2, 8, 0, 0, 0, 0, 3},
}
var formlist2 = Box_Num{
	{1, 2, 0, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 9, 1, 2, 3, 4, 5, 6},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{5, 6, 7, 8, 9, 1, 2, 3, 4},
	{8, 9, 1, 2, 3, 4, 5, 6, 7},
	{3, 4, 5, 6, 7, 8, 9, 1, 2},
	{6, 7, 8, 9, 1, 2, 3, 4, 5},
	{9, 1, 2, 3, 4, 5, 6, 7, 8},
}
var formlist3 = Box_Num{
	{1, 0, 9, 0, 0, 0, 0, 2, 0},
	{2, 0, 0, 5, 0, 0, 0, 0, 0},
	{3, 0, 6, 1, 0, 0, 8, 0, 0},
	{9, 0, 0, 0, 0, 3, 2, 0, 0},
	{0, 3, 5, 7, 0, 1, 4, 9, 0},
	{0, 0, 7, 9, 0, 0, 0, 0, 5},
	{0, 0, 1, 0, 0, 7, 9, 0, 2},
	{0, 0, 0, 0, 0, 5, 0, 0, 4},
	{0, 9, 0, 0, 0, 0, 3, 0, 8},
}

func main() {
	//fmt.Println(*form)
	/*//因为Form的定义是[9][9]Box，代表的是数组，虽然form是指针，
	//仍然可以直接使用form[i][j]取得数组中的值，也可以使用(*form)[i][j]，效果是一样的，
	//但如果Form是切片，则只能使用第二种方式
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			(*form)[i][j].Figure = Rand()
			(*form)[i][j].Lock = true
			fmt.Print((*form)[i][j].Figure, "  ")
		}
		fmt.Println("")
	}
	for i, v := range form {
		for j, _ := range v {
			fmt.Print(form[i][j].Figure, " ")
		}
		fmt.Println("")
		}
	*/
	var form = new(Form)
	form.Fill(Read())
	//form.Fill(formlist2)
	fmt.Println("已读取数独题：")
	form.Print(true)
	if Answer(form) {
		fmt.Println("成功解答，共计算了", Num+1, "次！")
		form.Print(true)
	}
}
func Read() Box_Num {
	var list Box_Num
	readbuf, _ := file.ReadFile(`E:/shudu/shudu.txt`)
	str := string(readbuf)
	n := 0
	for i, vlist := range list {
		for j, _ := range vlist {
			list[i][j] = uint8(str[n]) - 48
			n++
		}
	}
	return list
}
