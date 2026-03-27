package concurrent

import "sync"

// ConcurrentMap 泛型版本的并发安全映射实现
type ConcurrentMap[TKey comparable, TValue any] struct {
	mu  sync.RWMutex         // 读写锁，保证并发安全
	dic map[TKey]TValue      // 存储数据的映射
}

// NewConcurrentMap 创建一个新的泛型并发安全映射
// 参数:
//   Tkey: 键的类型，必须是可比较的
//   TValue: 值的类型
// 返回值:
//   *ConcurrentMap[Tkey, TValue]: 新创建的泛型并发安全映射
func NewConcurrentMap[Tkey comparable, TValue any]() *ConcurrentMap[Tkey, TValue] {
	return &ConcurrentMap[Tkey, TValue]{
		dic: make(map[Tkey]TValue),
	}
}

// Get 获取映射中指定键的值
// 参数:
//   key: 要查找的键
// 返回值:
//   TValue: 键对应的值
//   bool: 如果键存在，返回true；否则返回false
func (c *ConcurrentMap[TKey, TValue]) Get(key TKey) (TValue, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if v, ok := c.dic[key]; ok {
		return v, true
	} else {
		return v, false
	}
}

// Set 设置映射中指定键的值
// 参数:
//   key: 要设置的键
//   value: 要设置的值
func (c *ConcurrentMap[TKey, TValue]) Set(key TKey, value TValue) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dic[key] = value
}
