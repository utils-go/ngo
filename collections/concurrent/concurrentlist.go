package concurrent

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// changeType 表示列表变更类型
type changeType int

const (
	// add 表示添加操作
	add changeType = iota
	// remove 表示移除操作
	remove changeType = iota
)

// ConcurrentList 并发安全的列表实现
type ConcurrentList struct {
	data          []interface{}   // 存储数据的切片
	mux           sync.Mutex      // 互斥锁，保证并发安全
	chItemChanged chan changeType // 变更通知通道
}

// NewList 创建一个新的并发安全列表
// 参数:
//
//	T: 列表元素的类型
//
// 返回值:
//
//	*ConcurrentList: 新创建的并发安全列表
func NewList[T any]() *ConcurrentList {
	return &ConcurrentList{
		data:          make([]interface{}, 0),
		chItemChanged: make(chan changeType, 1),
	}
}

// notifyItemChanged 通知列表发生了变更
// 参数:
//
//	t: 变更类型
func (c *ConcurrentList) notifyItemChanged(t changeType) {
	select {
	case c.chItemChanged <- t:
	default:
	}
}

// Add 向列表中添加一个元素
// 参数:
//
//	v: 要添加的元素
func (c *ConcurrentList) Add(v interface{}) {
	c.mux.Lock()
	c.data = append(c.data, v)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

// AddRange 向列表中添加多个元素
// 参数:
//
//	v: 要添加的元素切片
func (c *ConcurrentList) AddRange(v []interface{}) {
	c.mux.Lock()
	c.data = append(c.data, v...)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

// Clear 清空列表
func (c *ConcurrentList) Clear() {
	c.mux.Lock()
	c.data = make([]interface{}, 0)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
}

// Remove 从列表中移除指定元素
// 参数:
//
//	v: 要移除的元素
//
// 返回值:
//
//	bool: 如果元素被成功移除，返回true；否则返回false
func (c *ConcurrentList) Remove(v interface{}) bool {
	c.mux.Lock()
	result := c.removeWithoutLock(v)
	c.mux.Unlock()

	c.notifyItemChanged(remove)

	return result
}

// removeWithoutLock 在无锁情况下从列表中移除指定元素
// 参数:
//
//	v: 要移除的元素
//
// 返回值:
//
//	bool: 如果元素被成功移除，返回true；否则返回false
func (c *ConcurrentList) removeWithoutLock(v interface{}) bool {
	newslice := make([]interface{}, 0, len(c.data))
	isexist := false

	for _, d := range c.data {
		if reflect.DeepEqual(d, v) && !isexist {
			isexist = true
		} else {
			newslice = append(newslice, d)
		}
	}

	if isexist {
		c.data = newslice
		return true
	}

	return false
}

// RemoveRange 从列表中移除指定范围的元素
// 参数:
//
//	index: 起始索引
//	count: 要移除的元素数量
func (c *ConcurrentList) RemoveRange(index, count int) {
	c.mux.Lock()
	c.removeRangeWithoutLock(index, count)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
}

// removeRangeWithoutLock 在无锁情况下从列表中移除指定范围的元素
// 参数:
//
//	index: 起始索引
//	count: 要移除的元素数量
func (c *ConcurrentList) removeRangeWithoutLock(index, count int) {
	newslice := make([]interface{}, 0, len(c.data)-1)
	newslice = append(newslice, c.data[0:index]...)
	newslice = append(newslice, c.data[index+count:len(c.data)]...)

	c.data = newslice
}

// Get 获取列表中指定索引的元素（不会移除）
// 参数:
//
//	index: 元素索引
//
// 返回值:
//
//	interface{}: 索引对应的元素
//	error: 如果索引越界，返回错误
func (c *ConcurrentList) Get(index int) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.getWithoutLock(index)
}

// getWithoutLock 在无锁情况下获取列表中指定索引的元素
// 参数:
//
//	index: 元素索引
//
// 返回值:
//
//	interface{}: 索引对应的元素
//	error: 如果索引越界，返回错误
func (c *ConcurrentList) getWithoutLock(index int) (interface{}, error) {
	var defaultT interface{}
	if len(c.data) <= index {
		return defaultT, errors.New(fmt.Sprintf("index: %d out of bound,max len: %d", index, len(c.data)))
	}
	return c.data[index], nil
}

// GetAll 获取列表中的所有元素（不会移除）
// 返回值:
//
//	[]interface{}: 列表中的所有元素
func (c *ConcurrentList) GetAll() []interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.data
}

// Take 获取并移除列表中指定索引的元素
// 参数:
//
//	index: 元素索引
//
// 返回值:
//
//	interface{}: 索引对应的元素
//	error: 如果索引越界或移除失败，返回错误
func (c *ConcurrentList) Take(index int) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	d, err := c.getWithoutLock(index)
	if err != nil {
		return d, err
	}
	if !c.removeWithoutLock(d) {
		return d, errors.New("remove fail")
	}
	return d, nil
}

// TakeAll 获取并移除列表中的所有元素
// 返回值:
//
//	[]interface{}: 列表中的所有元素
func (c *ConcurrentList) TakeAll() []interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()

	r := c.data
	c.data = make([]interface{}, 0)
	return r
}

// TakeAllBlock 阻塞直到列表中有元素添加，然后获取并移除所有元素
// 返回值:
//
//	[]interface{}: 列表中的所有元素
func (c *ConcurrentList) TakeAllBlock() []interface{} {
	for {
		if t := <-c.chItemChanged; t == add {
			break
		}
	}
	return c.TakeAll()
}
