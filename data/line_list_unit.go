package main

import "math"

const MAXSIZE int = 20
type ElemType int

type Sqlist struct {
	data [MAXSIZE]ElemType
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

func (l *Sqlist) clearList()  {
	if l.length == 0 {
		return
	}

	//for i := 0;i < l.length ; i++ {
	//	l.data[i] = ElemType(nil)
	//}

	l.length = 0
}

func (l *Sqlist) getElem(index int) ElemType {
	if index < 1 || index > l.length || l.length == 0 {
		return math.MaxInt64
	}
	return l.data[index - 1]
}

func (l *Sqlist) locateElem(v ElemType) int {
	for index, value := range l.data {
		if value == v {
			return index
			break
		}
	}
	return -1
}

func (l *Sqlist) listInsert(index int, v ElemType)  {
	if index >= MAXSIZE {
		return
	}
	if index < 1 || index > l.length + 1 {
		return
	}
	if index != l.length{
		for i := l.length ; i >= index ; i-- {
			l.data[i] = l.data[i - 1]
		}
	}
	l.data[index - 1] = v
	l.length++

}

func (l *Sqlist) listDelete(index int)  {
	if l.length == 0 {
		return
	}
	if index < 1 || index > l.length {
		return
	}
	if index < l.length {
		for i := l.length ; i >= index ; i-- {
			l.data[i] = l.data[i - 1]
		}
	}
	l.length--
}

func (l *Sqlist) listLength() int {
	return l.length
}
