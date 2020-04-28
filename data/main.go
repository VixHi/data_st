package main

import "fmt"

func main() {

	//初始化
	sqlist := new(Sqlist)

	//判断线性表是否为空
	fmt.Println(sqlist.listEmpty())

	//插入元素
	sqlist.listInsert(1, ElemType(2))
	sqlist.listInsert(2, ElemType(5))
	sqlist.listInsert(3, ElemType(8))

	sqlist.listInsert(1, ElemType(3))
	sqlist.listInsert(2, ElemType(6))
	sqlist.listInsert(3, ElemType(9))
	fmt.Println(sqlist.data, sqlist.length)

	//获取元素
	fmt.Println(sqlist.getElem(2))

	//删除元素
	sqlist.listDelete(4)
	fmt.Println(sqlist.data, sqlist.length)

	//查找元素 返回位置
	var elem ElemType = 5
	fmt.Println(sqlist.locateElem(elem))

	//清空线性表
	sqlist.clearList()
	fmt.Println(sqlist.data, sqlist.length)

	//判断线性表是否为空
	fmt.Println(sqlist.listEmpty())

}
