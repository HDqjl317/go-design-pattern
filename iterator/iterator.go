/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-19 08:45:47
 * @LastEditTime: 2023-04-19 08:55:54
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package iterator

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

type ArrayInt []int

// 数组迭代
type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

// 返回迭代器
func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}

func (iter *ArrayIntIterator) HasNext() bool {
	return iter.index < len(iter.arrayInt)-1
}

func (iter *ArrayIntIterator) Next() {
	iter.index++
}

func (iter *ArrayIntIterator) CurrentItem() interface{} {
	return iter.arrayInt[iter.index]
}
