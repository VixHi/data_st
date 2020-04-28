package main

import "math"

/*

线性表：顺序存储结构
	   优点 1.无需为表中元素之间的逻辑关系增加额外的存储空间 2.可以快速的存取表中任一位置的元素
	   缺点 1.插入和删除元素需要移动大量元素 2.线性表长度变化较大时难以确定存储空间的容量 3.造成空间的碎片
*/

// MAXSIZE 线性表最大存储
const MAXSIZE int = 20

// ElemType 元素类型
type ElemType int

// Sqlist 线性表
type Sqlist struct {
	data   [MAXSIZE]ElemType
	length int
}

func (l *Sqlist) initList() {
	var tempData [MAXSIZE]ElemType
	l.data = tempData
}

func (l *Sqlist) listEmpty() bool {
	if l.length == 0 {
		return true
	}
	return false
}

func (l *Sqlist) clearList() {
	if l.length == 0 {
		return
	}

	for i := 0; i < l.length; i++ {
		l.data[i] = ElemType(0)
	}

	l.length = 0
}

func (l *Sqlist) getElem(index int) ElemType {
	if index < 1 || index > l.length || l.length == 0 {
		return math.MaxInt64
	}
	return l.data[index-1]
}

func (l *Sqlist) locateElem(v ElemType) int {
	for index, value := range l.data {
		if value == v {
			return index + 1
			break
		}
	}
	return -1
}

func (l *Sqlist) listInsert(index int, v ElemType) {
	if index >= MAXSIZE {
		return
	}
	if index < 1 || index > l.length+1 {
		return
	}
	if index != l.length {
		for i := l.length; i >= index; i-- {
			l.data[i] = l.data[i-1]
		}
	}
	l.data[index-1] = v
	l.length++

}

func (l *Sqlist) listDelete(index int) {
	if l.length == 0 {
		return
	}
	if index < 1 || index > l.length {
		return
	}
	if index < l.length {
		for i := index; i < l.length; i++ {
			l.data[i-1] = l.data[i]
		}
	}
	l.data[l.length-1] = 0 //处理线性表最后一个元素
	l.length--
}

func (l *Sqlist) listLength() int {
	return l.length
}
