package concurrent

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// ConcurrentListT 泛型版本的并发安全列表实现
type ConcurrentListT[T any] struct {
	data          []T             // 存储数据的切片
	mux           sync.Mutex      // 互斥锁，保证并发安全
	chItemChanged chan changeType // 变更通知通道
}

// notifyItemChanged 通知列表发生了变更
// 参数:
//   t: 变更类型
func (c *ConcurrentListT[T]) notifyItemChanged(t changeType) {
	select {
	case c.chItemChanged <- t:
	default:
	}
}

// clearNotifyMsg 清空通知通道中的所有消息
func (c *ConcurrentListT[T]) clearNotifyMsg() {
	for {
		select {
		case _, ok := <-c.chItemChanged:
			if !ok {
				return
			}
		default:
			return
		}
	}
}

// NewListT 创建一个新的泛型并发安全列表
// 参数:
//   T: 列表元素的类型
// 返回值:
//   *ConcurrentListT[T]: 新创建的泛型并发安全列表
func NewListT[T any]() *ConcurrentListT[T] {
	return &ConcurrentListT[T]{
		data:          make([]T, 0),
		chItemChanged: make(chan changeType, 100),
	}
}

// Add 向列表中添加一个元素
// 参数:
//   v: 要添加的元素
func (c *ConcurrentListT[T]) Add(v T) {
	c.mux.Lock()
	c.data = append(c.data, v)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

// AddRange 向列表中添加多个元素
// 参数:
//   v: 要添加的元素切片
func (c *ConcurrentListT[T]) AddRange(v []T) {
	c.mux.Lock()
	c.data = append(c.data, v...)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

// Clear 清空列表
func (c *ConcurrentListT[T]) Clear() {
	c.mux.Lock()
	c.data = make([]T, 0)
	c.mux.Unlock()

	c.clearNotifyMsg()
	c.notifyItemChanged(remove)
}

// Remove 从列表中移除指定元素
// 参数:
//   v: 要移除的元素
// 返回值:
//   bool: 如果元素被成功移除，返回true；否则返回false
func (c *ConcurrentListT[T]) Remove(v T) bool {
	c.mux.Lock()
	result := c.removeWithoutLock(v)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
	return result
}

// removeWithoutLock 在无锁情况下从列表中移除指定元素
// 参数:
//   v: 要移除的元素
// 返回值:
//   bool: 如果元素被成功移除，返回true；否则返回false
func (c *ConcurrentListT[T]) removeWithoutLock(v T) bool {
	newslice := make([]T, 0)
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
//   index: 起始索引
//   count: 要移除的元素数量
func (c *ConcurrentListT[T]) RemoveRange(index, count int) {
	c.mux.Lock()
	c.removeRangeWithoutLock(index, count)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
}

// removeRangeWithoutLock 在无锁情况下从列表中移除指定范围的元素
// 参数:
//   index: 起始索引
//   count: 要移除的元素数量
func (c *ConcurrentListT[T]) removeRangeWithoutLock(index, count int) {
	newslice := make([]T, 0)
	newslice = append(newslice, c.data[0:index]...)
	newslice = append(newslice, c.data[index+count:len(c.data)]...)

	c.data = newslice
}

// Get 获取列表中指定索引的元素（不会移除）
// 参数:
//   index: 元素索引
// 返回值:
//   T: 索引对应的元素
//   error: 如果索引越界，返回错误
func (c *ConcurrentListT[T]) Get(index int) (T, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.getWithoutLock(index)
}

// getWithoutLock 在无锁情况下获取列表中指定索引的元素
// 参数:
//   index: 元素索引
// 返回值:
//   T: 索引对应的元素
//   error: 如果索引越界，返回错误
func (c *ConcurrentListT[T]) getWithoutLock(index int) (T, error) {
	var defaultT T
	if len(c.data) <= index {
		return defaultT, errors.New(fmt.Sprintf("index: %d out of bound,max len: %d", index, len(c.data)))
	}
	return c.data[index], nil
}

// GetAll 获取列表中的所有元素（不会移除）
// 返回值:
//   []T: 列表中的所有元素
func (c *ConcurrentListT[T]) GetAll() []T {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.data
}

// Take 获取并移除列表中指定索引的元素
// 参数:
//   index: 元素索引
// 返回值:
//   T: 索引对应的元素
//   error: 如果索引越界或移除失败，返回错误
func (c *ConcurrentListT[T]) Take(index int) (T, error) {
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
//   []T: 列表中的所有元素
func (c *ConcurrentListT[T]) TakeAll() []T {
	c.mux.Lock()
	defer c.mux.Unlock()

	r := c.data
	c.data = make([]T, 0)
	return r
}

// TakeAllBlock 阻塞直到列表中有元素添加或上下文被取消，然后获取并移除所有元素
// 参数:
//   ctx: 上下文，用于取消阻塞
// 返回值:
//   []T: 列表中的所有元素
//   bool: 如果成功获取元素，返回true；否则返回false
func (c *ConcurrentListT[T]) TakeAllBlock(ctx context.Context) ([]T, bool) {
	for {
		select {
		case t, ok := <-c.chItemChanged:
			if !ok {
				return nil, false
			}
			if t == add {
				return c.TakeAll(), true
			}
		case <-ctx.Done():
			return nil, false
		}
	}
}
