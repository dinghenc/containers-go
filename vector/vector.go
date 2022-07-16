package vector

import (
	"fmt"
)

const emptyVector = "vector is empty"

type vector struct {
	array []interface{}
}

// New 创建vector
func New() *vector {
	return &vector{}
}

// NewWithSize 使用初始大小创建vector
func NewWithSize(n int) *vector {
	return &vector{array: make([]interface{}, n)}
}

// Size 返回vector大小
func (v *vector) Size() int {
	return len(v.array)
}

// Empty vector是否为空
func (v *vector) Empty() bool {
	return v.Size() == 0
}

// Capacity 返回vector容量
func (v *vector) Capacity() int {
	return cap(v.array)
}

func (v *vector) checkOutOfRange(index int) bool {
	return index < 0 || index >= len(v.array)
}

// At 索引index的值
func (v *vector) At(index int) interface{} {
	if v.checkOutOfRange(index) {
		panic(outOfRange(v.Size(), index))
	}
	return v.array[index]
}

// Front 返回第一个元素
func (v *vector) Front() interface{} {
	if v.Empty() {
		panic(emptyVector)
	}
	return v.At(0)
}

// Back 返回最后一个元素
func (v *vector) Back() interface{} {
	if v.Empty() {
		panic(emptyVector)
	}
	return v.At(v.Size() - 1)
}

// Clear 清空vector内容
func (v *vector) Clear() {
	if v.Empty() {
		return
	}
	// 清空所有内容, 使得Size()为0, 但Capacity()不变
	v.array = make([]interface{}, 0, v.Capacity())
}

// Insert 在index前插入一个元素, 并返回当前的索引
func (v *vector) Insert(index int, val interface{}) int {
	if index == v.Size() {
		// 在末尾插入, 直接使用PushBack()
		v.PushBack(val)
		return index
	} else if v.checkOutOfRange(index) {
		panic(outOfRange(v.Size(), index))
	}
	// 先在末尾插入空值, 再从index位置处往后移动
	v.PushBack(nil)
	for i := v.Size() - 2; i >= index; i-- {
		v.array[i+1] = v.array[i]
	}
	v.array[index] = val
	return index
}

// Erase 移除指定位置的元素
func (v *vector) Erase(index int) interface{} {
	if v.checkOutOfRange(index) {
		panic(outOfRange(v.Size(), index))
	}
	// 将后面的元素往前挪动, 最后删除最后一个
	val := v.At(index)
	for i := index; i < v.Size()-1; i++ {
		v.array[i] = v.array[i+1]
	}
	v.PopBack()
	return val
}

// PushBack 从末尾追加元素
func (v *vector) PushBack(val interface{}) {
	v.array = append(v.array, val)
}

// PopBack 从末尾删除元素
func (v *vector) PopBack() {
	v.array = v.array[:v.Size()-1]
}

func outOfRange(size, index int) string {
	return fmt.Sprintf("out of range, length=%d, index=%d", size, index)
}
