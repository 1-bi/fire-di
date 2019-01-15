package fire_di

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const (
	taggedFuncStructFieldTag = "inject"
)

// FuncName returns a funcs formatted name
func funcName(fn interface{}) string {
	fnV := reflect.ValueOf(fn)
	if fnV.Kind() != reflect.Func {
		return "n/a"
	}

	function := runtime.FuncForPC(fnV.Pointer()).Name()
	return fmt.Sprintf("%s()", function)
}

func funcNameProvide(regBean *RegisterBean) string {

	var beanStr, funStr, funcName string
	beanTypeName := reflect.TypeOf(regBean.Bean)

	beanStr = beanTypeName.String()

	// --- check the func type
	funTypeName := reflect.ValueOf(regBean.ProvideFun).Elem()
	if funTypeName.Kind() != reflect.Func {
		funStr = "n/a"
	} else {
		funStr = funTypeName.Type().String()
	}

	funcName = strings.Join([]string{beanStr, funStr}, ":")
	return funcName
}

// FuncInterceptor define interceptor function handle
func FuncInterceptor(fptr interface{}, funIntercept func(in []reflect.Value) []reflect.Value) reflect.Value {
	fn := reflect.ValueOf(fptr).Elem()
	fn.Set(reflect.MakeFunc(fn.Type(), funIntercept))
	return fn
}

// whitelisting types to make sure the framework works
func isSupportedBindingKeyReflectType(reflectType reflect.Type) bool {
	return isSupportedBindReflectType(reflectType) || isSupportedBindInterfaceReflectType(reflectType) || isSupportedBindConstantReflectType(reflectType)
}

func isSupportedBindInterfaceReflectType(reflectType reflect.Type) bool {
	switch reflectType.Kind() {
	case reflect.Ptr:
		switch reflectType.Elem().Kind() {
		case reflect.Interface:
			return true
		default:
			return false
		}
	default:
		return false
	}
}

func isSupportedBindConstantReflectType(reflectType reflect.Type) bool {
	_, ok := constantKindForReflectType(reflectType)
	return ok
}

func isSupportedBindReflectType(reflectType reflect.Type) bool {
	switch reflectType.Kind() {
	case reflect.Ptr:
		switch reflectType.Elem().Kind() {
		case reflect.Interface:
			return true
		case reflect.Struct:
			return true
		default:
			return false
		}
	case reflect.Struct:
		return true
	default:
		return false
	}
}

func getStructFieldReflectTypeAndTag(structField reflect.StructField) (reflect.Type, string) {
	structFieldReflectType := structField.Type
	if structFieldReflectType.Kind() == reflect.Interface {
		structFieldReflectType = reflect.PtrTo(structFieldReflectType)
	}
	return structFieldReflectType, structField.Tag.Get(taggedFuncStructFieldTag)
}

func getTaggedFuncStructReflectValue(structReflectType reflect.Type, reflectValues []reflect.Value) *reflect.Value {
	structReflectValue := reflect.Indirect(reflect.New(structReflectType))
	populateStructReflectValue(&structReflectValue, reflectValues)
	return &structReflectValue
}

func newStructReflectValue(structReflectType reflect.Type) reflect.Value {
	return reflect.Indirect(reflect.New(structReflectType))
}

func populateStructReflectValue(structReflectValue *reflect.Value, reflectValues []reflect.Value) {
	numReflectValues := len(reflectValues)
	for i := 0; i < numReflectValues; i++ {
		structReflectValue.Field(i).Set(reflectValues[i])
	}
}

func isInterfacePtr(reflectType reflect.Type) bool {
	return isPtr(reflectType) && isInterface(reflectType.Elem())
}

func isStructPtr(reflectType reflect.Type) bool {
	return isPtr(reflectType) && isStruct(reflectType.Elem())
}

func isInterface(reflectType reflect.Type) bool {
	return reflectType.Kind() == reflect.Interface
}

func isStruct(reflectType reflect.Type) bool {
	return reflectType.Kind() == reflect.Struct
}

func isPtr(reflectType reflect.Type) bool {
	return reflectType.Kind() == reflect.Ptr
}

func isFunc(reflectType reflect.Type) bool {
	return reflectType.Kind() == reflect.Func
}
