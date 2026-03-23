package bitconverter

import (
	"bytes"
	"encoding/binary"
)

// BitConverter 提供在基本数据类型和字节数组之间进行转换的方法
// 参考: https://learn.microsoft.com/en-us/dotnet/api/system.bitconverter?view=netframework-4.7.2
type BitConverter struct {
	// IsLittleEndian 是否使用小端序
	IsLittleEndian bool
}

// toBytes 将任意类型转换为字节数组
// 参数:
//   value: 要转换的值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) toBytes(value any) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	var err error = nil
	if b.IsLittleEndian {
		err = binary.Write(buffer, binary.LittleEndian, value)
	} else {
		err = binary.Write(buffer, binary.BigEndian, value)
	}
	return buffer.Bytes(), err
}

// getValue 从字节数组中读取值到指定变量
// 参数:
//   value: 字节数组
//   v: 接收值的变量指针
// 返回值:
//   error: 读取过程中的错误
func (b *BitConverter) getValue(value []byte, v any) (err error) {
	buffer := bytes.NewBuffer(value)

	if b.IsLittleEndian {
		err = binary.Read(buffer, binary.LittleEndian, v)
	} else {
		err = binary.Read(buffer, binary.BigEndian, v)
	}
	return err
}

// GetBytesFromBoolE 将bool类型转换为字节数组
// 参数:
//   value: 要转换的bool值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromBoolE(value bool) ([]byte, error) {
	return b.toBytes(value)
}

// GetBytesFromInt16E 将int16类型转换为字节数组
// 参数:
//   value: 要转换的int16值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromInt16E(value int16) ([]byte, error) {
	return b.toBytes(value)
}

// GetBytesFromInt32E 将int32类型转换为字节数组
// 参数:
//   value: 要转换的int32值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromInt32E(value int32) ([]byte, error) {
	return b.toBytes(value)
}

// GetBytesFromInt64E 将int64类型转换为字节数组
// 参数:
//   value: 要转换的int64值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromInt64E(value int64) ([]byte, error) {
	return b.toBytes(value)
}

// GetBytesFromUInt16E 将uint16类型转换为字节数组
// 参数:
//   value: 要转换的uint16值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromUInt16E(value uint16) ([]byte, error) {
	return b.toBytes(value)
}

// GetBytesFromUInt32E 将uint32类型转换为字节数组
// 参数:
//   value: 要转换的uint32值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromUInt32E(value uint32) ([]byte, error) {
	return b.toBytes(value)
}

// GetBytesFromUInt64E 将uint64类型转换为字节数组
// 参数:
//   value: 要转换的uint64值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromUInt64E(value uint64) ([]byte, error) {
	return b.toBytes(value)
}

// GetBytesFromDoubleE 将float64类型转换为字节数组
// 参数:
//   value: 要转换的float64值
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func (b *BitConverter) GetBytesFromDoubleE(value float64) ([]byte, error) {
	return b.toBytes(value)
}

// ToBooleanE 将字节数组转换为bool类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   bool: 转换后的bool值
//   error: 转换过程中的错误
func (b *BitConverter) ToBooleanE(value []byte, startIndex int) (bool, error) {
	var result bool
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

// ToDoubleE 将字节数组转换为float64类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   float64: 转换后的float64值
//   error: 转换过程中的错误
func (b *BitConverter) ToDoubleE(value []byte, startIndex int) (float64, error) {
	var result float64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

// ToInt16E 将字节数组转换为int16类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   int16: 转换后的int16值
//   error: 转换过程中的错误
func (b *BitConverter) ToInt16E(value []byte, startIndex int) (int16, error) {
	var result int16
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

// ToInt32E 将字节数组转换为int32类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   int32: 转换后的int32值
//   error: 转换过程中的错误
func (b *BitConverter) ToInt32E(value []byte, startIndex int) (int32, error) {
	var result int32
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

// ToInt64E 将字节数组转换为int64类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   int64: 转换后的int64值
//   error: 转换过程中的错误
func (b *BitConverter) ToInt64E(value []byte, startIndex int) (int64, error) {
	var result int64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

// ToUInt16E 将字节数组转换为uint16类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   uint16: 转换后的uint16值
//   error: 转换过程中的错误
func (b *BitConverter) ToUInt16E(value []byte, startIndex int) (uint16, error) {
	var result uint16
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}
	return result, err
}

// ToUInt32E 将字节数组转换为uint32类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   uint32: 转换后的uint32值
//   error: 转换过程中的错误
func (b *BitConverter) ToUInt32E(value []byte, startIndex int) (uint32, error) {
	var result uint32
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

// ToUInt64E 将字节数组转换为uint64类型
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   uint64: 转换后的uint64值
//   error: 转换过程中的错误
func (b *BitConverter) ToUInt64E(value []byte, startIndex int) (uint64, error) {
	var result uint64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}
	return result, err
}

// GetBytesFromBool 将bool类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的bool值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromBool(value bool) []byte {
	v, _ := b.GetBytesFromBoolE(value)
	return v
}

// GetBytesFromInt16 将int16类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的int16值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromInt16(value int16) []byte {
	v, _ := b.GetBytesFromInt16E(value)
	return v
}

// GetBytesFromInt32 将int32类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的int32值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromInt32(value int32) []byte {
	v, _ := b.GetBytesFromInt32E(value)
	return v
}

// GetBytesFromInt64 将int64类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的int64值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromInt64(value int64) []byte {
	v, _ := b.toBytes(value)
	return v
}

// GetBytesFromUInt16 将uint16类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的uint16值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromUInt16(value uint16) []byte {
	v, _ := b.GetBytesFromUInt16E(value)
	return v
}

// GetBytesFromUInt32 将uint32类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的uint32值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromUInt32(value uint32) []byte {
	v, _ := b.GetBytesFromUInt32E(value)
	return v
}

// GetBytesFromUInt64 将uint64类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的uint64值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromUInt64(value uint64) []byte {
	v, _ := b.GetBytesFromUInt64E(value)
	return v
}

// GetBytesFromDouble 将float64类型转换为字节数组（忽略错误）
// 参数:
//   value: 要转换的float64值
// 返回值:
//   []byte: 转换后的字节数组
func (b *BitConverter) GetBytesFromDouble(value float64) []byte {
	v, _ := b.GetBytesFromDoubleE(value)
	return v
}

// ToBoolean 将字节数组转换为bool类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   bool: 转换后的bool值
func (b *BitConverter) ToBoolean(value []byte, startIndex int) bool {
	v, _ := b.ToBooleanE(value, startIndex)
	return v
}

// ToDouble 将字节数组转换为float64类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   float64: 转换后的float64值
func (b *BitConverter) ToDouble(value []byte, startIndex int) float64 {
	v, _ := b.ToDoubleE(value, startIndex)
	return v
}

// ToInt16 将字节数组转换为int16类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   int16: 转换后的int16值
func (b *BitConverter) ToInt16(value []byte, startIndex int) int16 {
	v, _ := b.ToInt16E(value, startIndex)
	return v
}

// ToInt32 将字节数组转换为int32类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   int32: 转换后的int32值
func (b *BitConverter) ToInt32(value []byte, startIndex int) int32 {
	v, _ := b.ToInt32E(value, startIndex)
	return v
}

// ToInt64 将字节数组转换为int64类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   int64: 转换后的int64值
func (b *BitConverter) ToInt64(value []byte, startIndex int) int64 {
	v, _ := b.ToInt64E(value, startIndex)
	return v
}

// ToUInt16 将字节数组转换为uint16类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   uint16: 转换后的uint16值
func (b *BitConverter) ToUInt16(value []byte, startIndex int) uint16 {
	v, _ := b.ToUInt16E(value, startIndex)
	return v
}

// ToUInt32 将字节数组转换为uint32类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   uint32: 转换后的uint32值
func (b *BitConverter) ToUInt32(value []byte, startIndex int) uint32 {
	v, _ := b.ToUInt32E(value, startIndex)
	return v
}

// ToUInt64 将字节数组转换为uint64类型（忽略错误）
// 参数:
//   value: 字节数组
//   startIndex: 开始转换的索引
// 返回值:
//   uint64: 转换后的uint64值
func (b *BitConverter) ToUInt64(value []byte, startIndex int) uint64 {
	v, _ := b.ToUInt64E(value, startIndex)
	return v
}
