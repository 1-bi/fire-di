package fire_di

import (
	"fmt"
	"gitlab.com/1-bi/log-api/loggercom"
	"reflect"
	"strings"
)

/**
 * define register bean
 */
type RegisterBean struct {
	// define base bean
	Bean       interface{}
	ProvideFun interface{}
}

/**
 * beanCtx for beanCtx
 */
type register struct {
	bindingErrors []error
	loginst       loggercom.Logger
	bindingFuns   map[string]interface{}
	bindingType   map[reflect.Type]reflect.Type
	invokedFuns   []interface{}
	beanFuns      map[string]interface{}
}

/**
 * register bean with the way "RegisterBean"
 */
func (this *register) RegBean(registerBean *RegisterBean) {

	// --- create new function ---
	proxyBean := this.getProxy(registerBean.Bean)

	proxyProvided := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		fn.Set(reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
			obj := reflect.ValueOf(registerBean.Bean)
			return []reflect.Value{obj}
		}))
	}

	proxyProvided(registerBean.ProvideFun)
	proxyHandler := reflect.ValueOf(registerBean.ProvideFun).Elem().Interface()

	fn := funcNameProvide(registerBean)
	this.bindingFuns[fn] = proxyHandler

	// --- check function with prefix "Inject" ---

	for methodName, m := range proxyBean.methods {

		var matchPrefix bool

		// --- check the use define prefix method prefix ---
		for _, prefix := range runnimeConf.injectMethodPrefix {

			if !strings.HasPrefix(methodName, prefix) {
				matchPrefix = true
				break
			}
		}

		if matchPrefix {
			// ---- invoke method bean ---
			this.Invoke(m.Func.Interface())

		}
	}

}

func (this *register) getProxy(ref interface{}) *proxyObject {
	proxyObj := new(proxyObject)
	proxyObj.ref = ref
	objType := reflect.TypeOf(ref)

	methodMap := make(map[string]reflect.Method, 0)
	var i int
	for i = 0; i < objType.NumMethod(); i++ {
		m := objType.Method(i)
		methodMap[m.Name] = m
	}
	proxyObj.methods = methodMap

	return proxyObj
}

// --- call and bind bean
func (this *register) InjectBean(beanFun interface{}) error {

	fn := funcName(beanFun)

	this.beanFuns[fn] = beanFun

	return nil
}

func (this *register) Invoke(handlers ...interface{}) error {
	// ---- get all handles ----
	for _, handler := range handlers {
		this.invokedFuns = append(this.invokedFuns, handler)
	}

	return nil
}

// ======================= private method ==================

/**
 * defined method binding
 * implement function for inject api
 */
func (this *register) String() string {
	return fmt.Sprintf("beanCtx{%s}", "update content ")
}

/**
 * create new beanCtx implement
 */
func createProvider() *register {
	return &register{make([]error, 0),
		nil, make(map[string]interface{}), make(map[reflect.Type]reflect.Type),
		make([]interface{}, 0), make(map[string]interface{}),
	}
}
