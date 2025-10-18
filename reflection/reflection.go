package reflection

import (
	"fmt"
	"reflect"
)

// Type represents type declarations equivalent to System.Type in .NET
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.type?view=netframework-4.7.2
type Type struct {
	goType reflect.Type
}

// MethodInfo represents information about a method
type MethodInfo struct {
	Name         string
	ReturnType   *Type
	Parameters   []*ParameterInfo
	IsPublic     bool
	IsStatic     bool
	DeclaringType *Type
	method       reflect.Method
	methodIndex  int
}



// ParameterInfo represents information about a parameter
type ParameterInfo struct {
	Name         string
	ParameterType *Type
	Position     int
	IsOptional   bool
}

// FieldInfo represents information about a field
type FieldInfo struct {
	Name         string
	FieldType    *Type
	IsPublic     bool
	IsStatic     bool
	DeclaringType *Type
	field        reflect.StructField
	fieldIndex   int
}

// ConstructorInfo represents information about a constructor
type ConstructorInfo struct {
	Parameters    []*ParameterInfo
	IsPublic      bool
	DeclaringType *Type
}

// MemberTypes represents the type of a member
type MemberTypes int

const (
	// Constructor represents a constructor member
	Constructor MemberTypes = iota
	// Event represents an event member
	Event
	// Field represents a field member
	Field
	// Method represents a method member
	Method
	// Property represents a property member
	Property
	// TypeInfo represents a type member
	TypeInfo
)

// BindingFlags specifies flags that control binding and the way in which the search for members is conducted
type BindingFlags int

const (
	// Default represents no binding flags
	Default BindingFlags = 0
	// Public includes public members in the search
	Public BindingFlags = 1 << iota
	// NonPublic includes non-public members in the search
	NonPublic
	// Instance includes instance members in the search
	Instance
	// Static includes static members in the search
	Static
	// FlattenHierarchy includes static members up the hierarchy
	FlattenHierarchy
)

// GetType returns the Type of the specified object
func GetType(obj interface{}) *Type {
	if obj == nil {
		return nil
	}
	
	goType := reflect.TypeOf(obj)
	return &Type{goType: goType}
}

// GetTypeFromName returns the Type with the specified name
func GetTypeFromName(typeName string) (*Type, error) {
	// This is a simplified implementation
	// In a full implementation, you would maintain a registry of types
	return nil, fmt.Errorf("GetTypeFromName not fully implemented for type: %s", typeName)
}

// Name gets the simple name of the current Type
func (t *Type) Name() string {
	if t == nil || t.goType == nil {
		return ""
	}
	
	// For pointer types, get the element type name
	if t.goType.Kind() == reflect.Ptr {
		return t.goType.Elem().Name()
	}
	
	return t.goType.Name()
}

// FullName gets the fully qualified name of the Type
func (t *Type) FullName() string {
	if t.goType == nil {
		return ""
	}
	
	pkg := t.goType.PkgPath()
	name := t.goType.Name()
	
	if pkg == "" {
		return name
	}
	return pkg + "." + name
}

// Namespace gets the namespace of the Type
func (t *Type) Namespace() string {
	if t.goType == nil {
		return ""
	}
	return t.goType.PkgPath()
}

// IsClass gets a value indicating whether the Type is a class
func (t *Type) IsClass() bool {
	if t.goType == nil {
		return false
	}
	return t.goType.Kind() == reflect.Struct
}

// IsInterface gets a value indicating whether the Type represents an interface
func (t *Type) IsInterface() bool {
	if t == nil || t.goType == nil {
		return false
	}
	
	// Check if it's a pointer to an interface
	if t.goType.Kind() == reflect.Ptr {
		return t.goType.Elem().Kind() == reflect.Interface
	}
	
	return t.goType.Kind() == reflect.Interface
}

// IsValueType gets a value indicating whether the Type is a value type
func (t *Type) IsValueType() bool {
	if t.goType == nil {
		return false
	}
	
	kind := t.goType.Kind()
	return kind != reflect.Ptr && kind != reflect.Interface && 
		   kind != reflect.Slice && kind != reflect.Map && 
		   kind != reflect.Chan && kind != reflect.Func
}

// IsPointer gets a value indicating whether the Type is a pointer
func (t *Type) IsPointer() bool {
	if t.goType == nil {
		return false
	}
	return t.goType.Kind() == reflect.Ptr
}

// IsArray gets a value indicating whether the Type is an array
func (t *Type) IsArray() bool {
	if t.goType == nil {
		return false
	}
	return t.goType.Kind() == reflect.Array
}

// IsSlice gets a value indicating whether the Type is a slice
func (t *Type) IsSlice() bool {
	if t.goType == nil {
		return false
	}
	return t.goType.Kind() == reflect.Slice
}

// GetMethods returns all methods of the current Type
func (t *Type) GetMethods() []*MethodInfo {
	return t.GetMethodsWithFlags(Public | Instance | Static)
}

// GetMethodsWithFlags returns methods that match the specified binding flags
func (t *Type) GetMethodsWithFlags(flags BindingFlags) []*MethodInfo {
	if t == nil || t.goType == nil {
		return nil
	}
	
	var methods []*MethodInfo
	
	// Get the actual type to work with (dereference pointers)
	actualType := t.goType
	if actualType.Kind() == reflect.Ptr {
		actualType = actualType.Elem()
	}
	
	// Get methods from the type
	for i := 0; i < t.goType.NumMethod(); i++ {
		method := t.goType.Method(i)
		
		// Check if method matches binding flags
		isPublic := isPublicName(method.Name)
		
		// If Public flag is set, only include public methods
		// If NonPublic flag is set, only include non-public methods
		// If both are set, include all methods
		if flags&Public != 0 && flags&NonPublic != 0 {
			// Include all methods
		} else if flags&Public != 0 && !isPublic {
			continue
		} else if flags&NonPublic != 0 && isPublic {
			continue
		} else if flags&Public == 0 && flags&NonPublic == 0 {
			// Default behavior: include public methods only
			if !isPublic {
				continue
			}
		}
		
		methodInfo := &MethodInfo{
			Name:          method.Name,
			ReturnType:    t.getReturnType(method.Type),
			Parameters:    t.getParameters(method.Type),
			IsPublic:      isPublic,
			IsStatic:      false, // Go methods are always instance methods
			DeclaringType: t,
			method:        method,
			methodIndex:   i,
		}
		
		methods = append(methods, methodInfo)
	}
	
	return methods
}

// GetMethod returns a method with the specified name
func (t *Type) GetMethod(name string) *MethodInfo {
	methods := t.GetMethods()
	for _, method := range methods {
		if method.Name == name {
			return method
		}
	}
	return nil
}



// GetFields returns all fields of the current Type
func (t *Type) GetFields() []*FieldInfo {
	return t.GetFieldsWithFlags(Public | Instance)
}

// GetFieldsWithFlags returns fields that match the specified binding flags
func (t *Type) GetFieldsWithFlags(flags BindingFlags) []*FieldInfo {
	if t == nil || t.goType == nil {
		return nil
	}
	
	// Get the actual struct type (dereference pointers)
	actualType := t.goType
	if actualType.Kind() == reflect.Ptr {
		actualType = actualType.Elem()
	}
	
	if actualType.Kind() != reflect.Struct {
		return nil
	}
	
	var fields []*FieldInfo
	
	for i := 0; i < actualType.NumField(); i++ {
		field := actualType.Field(i)
		
		// Check if field matches binding flags
		isPublic := isPublicName(field.Name)
		
		// If Public flag is set, only include public fields
		// If NonPublic flag is set, only include non-public fields
		// If both are set, include all fields
		if flags&Public != 0 && flags&NonPublic != 0 {
			// Include all fields
		} else if flags&Public != 0 && !isPublic {
			continue
		} else if flags&NonPublic != 0 && isPublic {
			continue
		} else if flags&Public == 0 && flags&NonPublic == 0 {
			// Default behavior: include public fields only
			if !isPublic {
				continue
			}
		}
		
		fieldInfo := &FieldInfo{
			Name:          field.Name,
			FieldType:     &Type{goType: field.Type},
			IsPublic:      isPublic,
			IsStatic:      false, // Go struct fields are not static
			DeclaringType: t,
			field:         field,
			fieldIndex:    i,
		}
		
		fields = append(fields, fieldInfo)
	}
	
	return fields
}

// GetField returns a field with the specified name
func (t *Type) GetField(name string) *FieldInfo {
	fields := t.GetFields()
	for _, field := range fields {
		if field.Name == name {
			return field
		}
	}
	return nil
}

// CreateInstance creates an instance of the Type
func (t *Type) CreateInstance(args ...interface{}) (interface{}, error) {
	if t.goType == nil {
		return nil, fmt.Errorf("cannot create instance of nil type")
	}
	
	// Handle pointer types
	targetType := t.goType
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}
	
	// Create new instance
	value := reflect.New(targetType)
	
	// If original type was not a pointer, return the element
	if t.goType.Kind() != reflect.Ptr {
		return value.Elem().Interface(), nil
	}
	
	return value.Interface(), nil
}

// IsAssignableFrom determines whether an instance of a specified Type can be assigned to the current Type
func (t *Type) IsAssignableFrom(other *Type) bool {
	if t.goType == nil || other == nil || other.goType == nil {
		return false
	}
	
	return other.goType.AssignableTo(t.goType)
}

// Invoke calls the method represented by the current MethodInfo
func (m *MethodInfo) Invoke(obj interface{}, args ...interface{}) ([]interface{}, error) {
	if obj == nil {
		return nil, fmt.Errorf("cannot invoke method on nil object")
	}
	
	objValue := reflect.ValueOf(obj)
	method := objValue.Method(m.methodIndex)
	
	// Convert arguments to reflect.Value
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}
	
	// Call the method
	results := method.Call(argValues)
	
	// Convert results to interface{}
	resultInterfaces := make([]interface{}, len(results))
	for i, result := range results {
		resultInterfaces[i] = result.Interface()
	}
	
	return resultInterfaces, nil
}



// GetValue gets the value of the field for the specified object
func (f *FieldInfo) GetValue(obj interface{}) (interface{}, error) {
	if obj == nil {
		return nil, fmt.Errorf("cannot get field value from nil object")
	}
	
	objValue := reflect.ValueOf(obj)
	
	// Handle pointer types
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	
	if objValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("object is not a struct")
	}
	
	fieldValue := objValue.Field(f.fieldIndex)
	return fieldValue.Interface(), nil
}

// SetValue sets the value of the field for the specified object
func (f *FieldInfo) SetValue(obj interface{}, value interface{}) error {
	if obj == nil {
		return fmt.Errorf("cannot set field value on nil object")
	}
	
	objValue := reflect.ValueOf(obj)
	
	// Handle pointer types
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	
	if objValue.Kind() != reflect.Struct {
		return fmt.Errorf("object is not a struct")
	}
	
	fieldValue := objValue.Field(f.fieldIndex)
	
	if !fieldValue.CanSet() {
		return fmt.Errorf("field %s cannot be set", f.Name)
	}
	
	newValue := reflect.ValueOf(value)
	if !newValue.Type().AssignableTo(fieldValue.Type()) {
		return fmt.Errorf("cannot assign %T to field of type %s", value, fieldValue.Type())
	}
	
	fieldValue.Set(newValue)
	return nil
}

// Helper functions

func isPublicName(name string) bool {
	if len(name) == 0 {
		return false
	}
	return name[0] >= 'A' && name[0] <= 'Z'
}

func (t *Type) getReturnType(methodType reflect.Type) *Type {
	if methodType.NumOut() == 0 {
		return nil
	}
	
	// Return the first return type (ignoring error returns for simplicity)
	return &Type{goType: methodType.Out(0)}
}

func (t *Type) getParameters(methodType reflect.Type) []*ParameterInfo {
	var params []*ParameterInfo
	
	// Skip the first parameter (receiver) for methods
	start := 1
	if methodType.NumIn() > 0 {
		for i := start; i < methodType.NumIn(); i++ {
			paramType := methodType.In(i)
			param := &ParameterInfo{
				Name:          fmt.Sprintf("param%d", i-start),
				ParameterType: &Type{goType: paramType},
				Position:      i - start,
				IsOptional:    false,
			}
			params = append(params, param)
		}
	}
	
	return params
}

// String returns a string representation of the Type
func (t *Type) String() string {
	if t == nil || t.goType == nil {
		return "<nil>"
	}
	return t.goType.String()
}

// Equals determines whether the current Type is equal to another Type
func (t *Type) Equals(other *Type) bool {
	if t == nil && other == nil {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	return t.goType == other.goType
}

// GetElementType returns the Type of the object encompassed or referred to by the current array, pointer, or reference type
func (t *Type) GetElementType() *Type {
	if t == nil || t.goType == nil {
		return nil
	}
	
	kind := t.goType.Kind()
	if kind == reflect.Array || kind == reflect.Slice || kind == reflect.Ptr {
		return &Type{goType: t.goType.Elem()}
	}
	
	return nil
}

// GetInterfaces returns all interfaces implemented by the current Type
func (t *Type) GetInterfaces() []*Type {
	if t.goType == nil {
		return nil
	}
	
	var interfaces []*Type
	
	for i := 0; i < t.goType.NumMethod(); i++ {
		// This is a simplified implementation
		// In Go, we can't easily get all implemented interfaces
		// This would require a more complex type registry system
	}
	
	return interfaces
}