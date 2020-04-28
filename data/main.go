package main

import "fmt"

func main() {

sqlist := new(Sqlist)
sqlist.listInsert(0, ElemType(2))
sqlist.listInsert(0, ElemType(0))
sqlist.listInsert(0, ElemType(8))
fmt.Println(sqlist.data, sqlist.length)
}

