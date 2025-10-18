package reflection

import (
	"testing"
)

// Test structures and types
type TestStruct struct {
	PublicField    string
	privateField   int
	AnotherField   bool
	NumericField   float64
}

func (ts *TestStruct) PublicMethod(arg string) string {
	return "Hello, " + arg
}

func (ts *TestStruct) privateMethod() int {
	return 42
}

func (ts *TestStruct) MethodWithMultipleParams(a int, b string) (int, string) {
	return a * 2, b + " processed"
}

func (ts *TestStruct) MethodWithNoReturn(value int) {
	ts.privateField = value
}

type TestInterface interface {
	InterfaceMethod() string
}

func (ts *TestStruct) InterfaceMethod() string {
	return "interface implementation"
}

func TestGetType(t *testing.T) {
	obj := &TestStruct{
		PublicField:  "test",
		privateField: 123,
		AnotherField: true,
		NumericField: 3.14,
	}
	
	typ := GetType(obj)
	
	if typ == nil {
		t.Fatal("GetType returned nil")
	}
	
	if typ.Name() != "TestStruct" {
		t.Errorf("Expected type name 'TestStruct', got %q", typ.Name())
	}
	
	if !typ.IsPointer() {
		t.Error("Expected type to be a pointer")
	}
	
	// Test with non-pointer
	obj2 := TestStruct{}
	typ2 := GetType(obj2)
	
	if typ2.IsPointer() {
		t.Error("Expected type to not be a pointer")
	}
	
	if !typ2.IsClass() {
		t.Error("Expected type to be a class (struct)")
	}
}

func TestGetTypeNil(t *testing.T) {
	typ := GetType(nil)
	if typ != nil {
		t.Error("GetType(nil) should return nil")
	}
}

func TestTypeProperties(t *testing.T) {
	obj := TestStruct{}
	typ := GetType(obj)
	
	// Test basic properties
	if typ.Name() != "TestStruct" {
		t.Errorf("Expected Name() = 'TestStruct', got %q", typ.Name())
	}
	
	expectedFullName := "github.com/utils-go/ngo/reflection.TestStruct"
	if typ.FullName() != expectedFullName {
		t.Errorf("Expected FullName() = %q, got %q", expectedFullName, typ.FullName())
	}
	
	expectedNamespace := "github.com/utils-go/ngo/reflection"
	if typ.Namespace() != expectedNamespace {
		t.Errorf("Expected Namespace() = %q, got %q", expectedNamespace, typ.Namespace())
	}
	
	if !typ.IsClass() {
		t.Error("Expected IsClass() = true")
	}
	
	if typ.IsInterface() {
		t.Error("Expected IsInterface() = false")
	}
	
	if !typ.IsValueType() {
		t.Error("Expected IsValueType() = true")
	}
	
	if typ.IsPointer() {
		t.Error("Expected IsPointer() = false")
	}
}

func TestGetMethods(t *testing.T) {
	obj := &TestStruct{}
	typ := GetType(obj)
	
	methods := typ.GetMethods()
	
	if len(methods) == 0 {
		t.Error("Expected to find methods")
	}
	
	// Find specific methods
	var publicMethod *MethodInfo
	var multiParamMethod *MethodInfo
	
	for _, method := range methods {
		switch method.Name {
		case "PublicMethod":
			publicMethod = method
		case "MethodWithMultipleParams":
			multiParamMethod = method
		}
	}
	
	if publicMethod == nil {
		t.Error("PublicMethod not found")
	} else {
		if !publicMethod.IsPublic {
			t.Error("PublicMethod should be public")
		}
		
		if publicMethod.ReturnType == nil {
			t.Error("PublicMethod should have a return type")
		}
		
		if len(publicMethod.Parameters) != 1 {
			t.Errorf("PublicMethod should have 1 parameter, got %d", len(publicMethod.Parameters))
		}
	}
	
	if multiParamMethod == nil {
		t.Error("MethodWithMultipleParams not found")
	} else {
		if len(multiParamMethod.Parameters) != 2 {
			t.Errorf("MethodWithMultipleParams should have 2 parameters, got %d", len(multiParamMethod.Parameters))
		}
	}
}

func TestGetMethod(t *testing.T) {
	obj := &TestStruct{}
	typ := GetType(obj)
	
	method := typ.GetMethod("PublicMethod")
	if method == nil {
		t.Fatal("GetMethod('PublicMethod') returned nil")
	}
	
	if method.Name != "PublicMethod" {
		t.Errorf("Expected method name 'PublicMethod', got %q", method.Name)
	}
	
	// Test non-existent method
	nonExistent := typ.GetMethod("NonExistentMethod")
	if nonExistent != nil {
		t.Error("GetMethod should return nil for non-existent method")
	}
}





func TestGetFields(t *testing.T) {
	obj := TestStruct{}
	typ := GetType(obj)
	
	fields := typ.GetFields()
	
	expectedPublicFields := []string{"PublicField", "AnotherField", "NumericField"}
	foundFields := make(map[string]bool)
	
	for _, field := range fields {
		foundFields[field.Name] = true
		
		if !field.IsPublic {
			t.Errorf("Field %s should be public", field.Name)
		}
	}
	
	for _, expected := range expectedPublicFields {
		if !foundFields[expected] {
			t.Errorf("Expected to find public field %s", expected)
		}
	}
}

func TestGetField(t *testing.T) {
	obj := TestStruct{}
	typ := GetType(obj)
	
	field := typ.GetField("PublicField")
	if field == nil {
		t.Fatal("GetField('PublicField') returned nil")
	}
	
	if field.Name != "PublicField" {
		t.Errorf("Expected field name 'PublicField', got %q", field.Name)
	}
	
	if field.FieldType.Name() != "string" {
		t.Errorf("Expected field type 'string', got %q", field.FieldType.Name())
	}
}

func TestCreateInstance(t *testing.T) {
	typ := GetType(TestStruct{})
	
	instance, err := typ.CreateInstance()
	if err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}
	
	if _, ok := instance.(TestStruct); !ok {
		t.Errorf("Expected instance to be TestStruct, got %T", instance)
	}
	
	// Test with pointer type
	ptrType := GetType(&TestStruct{})
	ptrInstance, err := ptrType.CreateInstance()
	if err != nil {
		t.Fatalf("CreateInstance for pointer type failed: %v", err)
	}
	
	if _, ok := ptrInstance.(*TestStruct); !ok {
		t.Errorf("Expected pointer instance to be *TestStruct, got %T", ptrInstance)
	}
	
	_ = ptrType // Use the variable to avoid "declared and not used" error
}

func TestMethodInvoke(t *testing.T) {
	obj := &TestStruct{
		PublicField: "test",
	}
	
	typ := GetType(obj)
	method := typ.GetMethod("PublicMethod")
	
	if method == nil {
		t.Fatal("Could not find PublicMethod")
	}
	
	results, err := method.Invoke(obj, "World")
	if err != nil {
		t.Fatalf("Method invocation failed: %v", err)
	}
	
	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}
	
	result, ok := results[0].(string)
	if !ok {
		t.Fatalf("Expected string result, got %T", results[0])
	}
	
	expected := "Hello, World"
	if result != expected {
		t.Errorf("Expected result %q, got %q", expected, result)
	}
}



func TestFieldGetSetValue(t *testing.T) {
	obj := &TestStruct{
		AnotherField: false,
	}
	
	typ := GetType(obj)
	field := typ.GetField("AnotherField")
	
	if field == nil {
		t.Fatal("Could not find AnotherField")
	}
	
	// Test GetValue
	value, err := field.GetValue(obj)
	if err != nil {
		t.Fatalf("GetValue failed: %v", err)
	}
	
	if value != false {
		t.Errorf("Expected value false, got %v", value)
	}
	
	// Test SetValue
	err = field.SetValue(obj, true)
	if err != nil {
		t.Fatalf("SetValue failed: %v", err)
	}
	
	// Verify the change
	if obj.AnotherField != true {
		t.Errorf("Expected field to be true, got %v", obj.AnotherField)
	}
}

func TestIsAssignableFrom(t *testing.T) {
	structType := GetType(TestStruct{})
	ptrType := GetType(&TestStruct{})
	stringType := GetType("")
	
	// Test same types
	if !structType.IsAssignableFrom(structType) {
		t.Error("Type should be assignable from itself")
	}
	
	// Test different types
	if structType.IsAssignableFrom(stringType) {
		t.Error("TestStruct should not be assignable from string")
	}
	
	// Test pointer vs struct (should not be assignable)
	if structType.IsAssignableFrom(ptrType) {
		t.Error("TestStruct should not be assignable from *TestStruct")
	}
	
	// Test nil
	if structType.IsAssignableFrom(nil) {
		t.Error("Should not be assignable from nil type")
	}
}

func TestTypeEquals(t *testing.T) {
	type1 := GetType(TestStruct{})
	type2 := GetType(TestStruct{})
	type3 := GetType("")
	
	if !type1.Equals(type2) {
		t.Error("Same types should be equal")
	}
	
	if type1.Equals(type3) {
		t.Error("Different types should not be equal")
	}
	
	if type1.Equals(nil) {
		t.Error("Type should not equal nil")
	}
	
	var nilType *Type
	if !nilType.Equals(nil) {
		t.Error("Nil type should equal nil")
	}
}

func TestGetElementType(t *testing.T) {
	// Test array
	arrayType := GetType([5]int{})
	elementType := arrayType.GetElementType()
	
	if elementType == nil {
		t.Fatal("Array element type should not be nil")
	}
	
	if elementType.Name() != "int" {
		t.Errorf("Expected element type 'int', got %q", elementType.Name())
	}
	
	// Test slice
	sliceType := GetType([]string{})
	elementType = sliceType.GetElementType()
	
	if elementType == nil {
		t.Fatal("Slice element type should not be nil")
	}
	
	if elementType.Name() != "string" {
		t.Errorf("Expected element type 'string', got %q", elementType.Name())
	}
	
	// Test pointer
	ptrType := GetType(&TestStruct{})
	elementType = ptrType.GetElementType()
	
	if elementType == nil {
		t.Fatal("Pointer element type should not be nil")
	}
	
	if elementType.Name() != "TestStruct" {
		t.Errorf("Expected element type 'TestStruct', got %q", elementType.Name())
	}
	
	// Test non-element type
	structType := GetType(TestStruct{})
	elementType = structType.GetElementType()
	
	if elementType != nil {
		t.Error("Struct should not have element type")
	}
}

func TestTypeKindChecks(t *testing.T) {
	tests := []struct {
		obj        interface{}
		isClass    bool
		isInterface bool
		isValueType bool
		isPointer  bool
		isArray    bool
		isSlice    bool
	}{
		{TestStruct{}, true, false, true, false, false, false},
		{&TestStruct{}, false, false, false, true, false, false},
		{(*TestInterface)(nil), false, true, false, true, false, false},
		{[5]int{}, false, false, true, false, true, false},
		{[]int{}, false, false, false, false, false, true},
		{"string", false, false, true, false, false, false},
		{42, false, false, true, false, false, false},
	}
	
	for i, test := range tests {
		typ := GetType(test.obj)
		
		if typ.IsClass() != test.isClass {
			t.Errorf("Test %d: IsClass() = %v, expected %v", i, typ.IsClass(), test.isClass)
		}
		
		if typ.IsInterface() != test.isInterface {
			t.Errorf("Test %d: IsInterface() = %v, expected %v", i, typ.IsInterface(), test.isInterface)
		}
		
		if typ.IsValueType() != test.isValueType {
			t.Errorf("Test %d: IsValueType() = %v, expected %v", i, typ.IsValueType(), test.isValueType)
		}
		
		if typ.IsPointer() != test.isPointer {
			t.Errorf("Test %d: IsPointer() = %v, expected %v", i, typ.IsPointer(), test.isPointer)
		}
		
		if typ.IsArray() != test.isArray {
			t.Errorf("Test %d: IsArray() = %v, expected %v", i, typ.IsArray(), test.isArray)
		}
		
		if typ.IsSlice() != test.isSlice {
			t.Errorf("Test %d: IsSlice() = %v, expected %v", i, typ.IsSlice(), test.isSlice)
		}
	}
}

func TestBindingFlags(t *testing.T) {
	obj := &TestStruct{}
	typ := GetType(obj)
	
	// Test public only
	publicMethods := typ.GetMethodsWithFlags(Public | Instance)
	publicFields := typ.GetFieldsWithFlags(Public | Instance)
	
	// Should find public members
	if len(publicMethods) == 0 {
		t.Error("Should find public methods")
	}
	
	if len(publicFields) == 0 {
		t.Error("Should find public fields")
	}
	
	// Verify all found members are public
	for _, method := range publicMethods {
		if !method.IsPublic {
			t.Errorf("Method %s should be public", method.Name)
		}
	}
	
	for _, field := range publicFields {
		if !field.IsPublic {
			t.Errorf("Field %s should be public", field.Name)
		}
	}
}

func TestTypeString(t *testing.T) {
	typ := GetType(TestStruct{})
	str := typ.String()
	
	expected := "reflection.TestStruct"
	if str != expected {
		t.Errorf("Expected String() = %q, got %q", expected, str)
	}
	
	// Test nil type
	var nilType *Type
	if nilType.String() != "<nil>" {
		t.Errorf("Expected nil type String() = '<nil>', got %q", nilType.String())
	}
}

func BenchmarkGetType(b *testing.B) {
	obj := &TestStruct{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetType(obj)
	}
}

func BenchmarkGetMethods(b *testing.B) {
	obj := &TestStruct{}
	typ := GetType(obj)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		typ.GetMethods()
	}
}

func BenchmarkMethodInvoke(b *testing.B) {
	obj := &TestStruct{}
	typ := GetType(obj)
	method := typ.GetMethod("PublicMethod")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		method.Invoke(obj, "test")
	}
}

