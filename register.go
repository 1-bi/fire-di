package fire_di

import (
	"errors"
	"fmt"
	"github.com/1-bi/log-api"
	"reflect"
	"strings"
)

// RegisterBean define register bean
type RegisterBean struct {
	// define base bean
	Bean       interface{}
	ProvideFun interface{}
}

/**
 * create new beanCtx implement
 */
func newRegister() *register {
	register := new(register)
	register.bindingErrors = make([]error, 0)
	register.beanFuns = make(map[string]interface{})
	register.bindingFuns = make(map[string]interface{})
	register.bindingType = make(map[reflect.Type]reflect.Type)
	register.invokedFuns = make([]interface{}, 0)
	register.loginst = nil

	register.proxyBeans = make([]*BeanProxy, 0)

	return register
}

/**
 * bean register for bindding mapping --
 */
type register struct {
	bindingErrors []error
	loginst       logapi.Logger
	bindingFuns   map[string]interface{}
	bindingType   map[reflect.Type]reflect.Type
	invokedFuns   []interface{}
	beanFuns      map[string]interface{}
	// define proxy bean dependency
	proxyBeans []*BeanProxy
}

func (myself *register) convertToResultObject(registerBean *RegisterBean) (reflect.Value, error) {
	beanVal := reflect.ValueOf(registerBean.Bean)
	funVal := reflect.ValueOf(registerBean.ProvideFun)
	funTyp := funVal.Type().Elem()

	var outputObj reflect.Value

	if funTyp.NumOut() != 1 {
		return outputObj, errors.New("The number of return object is not equal to 1 . ")
	}

	returnOutTyp := funTyp.Out(0)

	if returnOutTyp.Kind() == reflect.Interface {

		if beanVal.Type().Implements(returnOutTyp) {
			outputObj = beanVal.Convert(returnOutTyp)
		}
	} else if returnOutTyp.Kind() == reflect.Ptr {
		outputObj = beanVal
	} else if returnOutTyp.Kind() == reflect.Struct {

		instTyp := beanVal.Type().String()
		retoutTyp := returnOutTyp.String()

		if strings.HasSuffix(instTyp, retoutTyp) {
			outputObj = beanVal.Elem()
		} else {
			return outputObj, errors.New("Implement \"" + beanVal.Type().String() + "\" is not match struct \"" + returnOutTyp.String() + "\"")
		}
	}

	return outputObj, nil

}

/**
 * register bean with the way "RegisterBean"
 */
func (myself *register) RegBean(registerBean *RegisterBean) {

	// build inject object bean method
	proxyBean := myself.getProxy(registerBean.Bean)

	outputObj, err := myself.convertToResultObject(registerBean)

	if err != nil {
		myself.loginst.Info(err.Error(), nil)

	}

	proxyHandlerRef := FuncInterceptor(registerBean.ProvideFun, func(in []reflect.Value) []reflect.Value {

		return []reflect.Value{outputObj}
	})

	proxyHandler := proxyHandlerRef.Interface()

	fn := funcNameProvide(registerBean)
	myself.bindingFuns[fn] = proxyHandler

	// --- append the proxy bean ---
	myself.proxyBeans = append(myself.proxyBeans, proxyBean)

	// --- check function with prefix "Inject" ---

}

// RegFunc set the function for register function
func (myself *register) RegFunc(fn interface{}) {

	// --- defined function mapping ---
	fnPrt := reflect.ValueOf(fn)

	newFunType := reflect.New(fnPrt.Type())

	resultFun := FuncInterceptor(newFunType.Interface(), func(in []reflect.Value) []reflect.Value {

		// --- defined depenMethods status ---
		result := fnPrt.Call(in)

		return result
	})

	// define object ---
	/*
		var params = make([]reflect.Value, fnPrt.Type().NumIn())

		mockObj := new(mockobject.Case4MockObj2)
		params[0] = reflect.ValueOf(mockObj)
	*/

	fName := funcName(fn)
	myself.bindingFuns[fName] = resultFun.Interface()

}

// GetProxyBeans get the proxy beans reference
func (myself *register) GetProxyBeans() []*BeanProxy {
	return myself.proxyBeans
}

func (myself *register) getProxy(ref interface{}) *BeanProxy {
	proxyObj := new(BeanProxy)
	proxyObj.dependentStructs = make([]string, 0)
	proxyObj.applyProxy(ref)
	return proxyObj
}

// --- call and bind bean
func (myself *register) InjectBean(beanFun interface{}) error {

	fn := funcName(beanFun)

	myself.beanFuns[fn] = beanFun

	return nil
}

func (myself *register) Invoke(handlers ...interface{}) error {
	// ---- get all handles ----
	for _, handler := range handlers {
		myself.invokedFuns = append(myself.invokedFuns, handler)
	}

	return nil
}

// ======================= private method ==================

/**
 * defined method binding
 * implement function for inject api
 */
func (myself *register) String() string {
	return fmt.Sprintf("beanCtx{%s}", "update content ")
}
