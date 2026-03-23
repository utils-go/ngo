package reflection

import (
	"testing"
)

func TestAssembly_GetExecutingAssembly(t *testing.T) {
	assembly := GetExecutingAssembly()

	// 检查返回的Assembly是否非空
	if assembly == nil {
		t.Error("GetExecutingAssembly() should return a non-nil Assembly")
	}
}

func TestActivator_CreateInstance(t *testing.T) {
	// 获取TestStruct的类型
	testInstance := &TestStruct{}
	testType := GetType(testInstance)

	// 创建Activator
	activator := NewActivator()

	// 使用Activator创建实例
	newInstance, err := activator.CreateInstance(testType)
	if err != nil {
		t.Errorf("Activator.CreateInstance() failed: %v", err)
	}

	// 检查创建的实例是否非空
	if newInstance == nil {
		t.Error("Activator.CreateInstance() should return a non-nil instance")
	}

	// 检查创建的实例类型是否正确
	if _, ok := newInstance.(*TestStruct); !ok {
		t.Error("Activator.CreateInstance() should return an instance of TestStruct")
	}
}

func TestActivator_CreateInstanceWithArgs(t *testing.T) {
	// 测试带参数的构造函数
	// 注意：Go的struct没有构造函数，所以这个测试主要验证基本功能
	testInstance := &TestStruct{}
	testType := GetType(testInstance)

	// 创建Activator
	activator := NewActivator()

	// 尝试使用参数创建实例
	// 虽然Go的struct没有构造函数，但CreateInstance方法应该能够处理这种情况
	newInstance, err := activator.CreateInstance(testType, "test")
	if err != nil {
		t.Errorf("Activator.CreateInstance() with args failed: %v", err)
	}

	// 检查创建的实例是否非空
	if newInstance == nil {
		t.Error("Activator.CreateInstance() should return a non-nil instance")
	}
}
