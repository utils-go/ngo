package concurrent

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type ConcurrentListT[T any] struct {
	data          []T
	mux           sync.Mutex
	chItemChanged chan changeType
}

func (c *ConcurrentListT[T]) notifyItemChanged(t changeType) {
	select {
	case c.chItemChanged <- t:
	default:
	}
}

func NewListT[T any]() *ConcurrentListT[T] {
	return &ConcurrentListT[T]{
		data:          make([]T, 0),
		chItemChanged: make(chan changeType, 1),
	}
}

func (c *ConcurrentListT[T]) Add(v T) {
	c.mux.Lock()
	c.data = append(c.data, v)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

func (c *ConcurrentListT[T]) AddRange(v []T) {
	c.mux.Lock()
	c.data = append(c.data, v...)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

func (c *ConcurrentListT[T]) Clear() {
	c.mux.Lock()
	c.data = make([]T, 0)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
}

func (c *ConcurrentListT[T]) Remove(v T) bool {
	c.mux.Lock()
	result := c.removeWithoutLock(v)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
	return result
}
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

func (c *ConcurrentListT[T]) RemoveRange(index, count int) {
	c.mux.Lock()
	c.removeRangeWithoutLock(index, count)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
}
func (c *ConcurrentListT[T]) removeRangeWithoutLock(index, count int) {
	newslice := make([]T, 0)
	newslice = append(newslice, c.data[0:index]...)
	newslice = append(newslice, c.data[index+count:len(c.data)]...)

	c.data = newslice
}

// just get ,not remove
func (c *ConcurrentListT[T]) Get(index int) (T, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.getWithoutLock(index)
}

func (c *ConcurrentListT[T]) getWithoutLock(index int) (T, error) {
	var defaultT T
	if len(c.data) <= index {
		return defaultT, errors.New(fmt.Sprintf("index: %d out of bound,max len: %d", index, len(c.data)))
	}
	return c.data[index], nil
}

// get all and not remove
func (c *ConcurrentListT[T]) GetAll() []T {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.data
}

// get one item and remove
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
func (c *ConcurrentListT[T]) TakeAll() []T {
	c.mux.Lock()
	defer c.mux.Unlock()

	r := c.data
	c.data = make([]T, 0)
	return r
}

// when no item add or no cancel,it will block
func (c *ConcurrentListT[T]) TakeAllBlock(ctx context.Context) (bool, []T) {
	chHasAdd := make(chan struct{})
	go func() {
		for {
			if t := <-c.chItemChanged; t == add {
				chHasAdd <- struct{}{}
			}
		}
	}()

	select {
	case <-ctx.Done():
		return false, nil
	case <-chHasAdd:
		return true, c.TakeAll()
	}
}
