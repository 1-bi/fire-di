package fire_di

import (
	"fmt"
	"gitlab.com/1-bi/log-api/loggercom"
	"reflect"
)

/**
 * beanCtx for beanCtx
 */
type provider struct {
	bindingErrors []error
	loginst       loggercom.Logger
	bindingFuns   map[string]interface{}
	bindingType   map[reflect.Type]reflect.Type
	invokedFuns   []interface{}
	beanFuns      map[string]interface{}
}

/**
 * @deprecate method
 */
func (this *provider) Provide(handlers ...interface{}) error {

	// ---- get all handles ----
	for _, handler := range handlers {

		fn := funcName(handler)

		this.bindingFuns[fn] = handler

	}
	return nil
}

/**
 * add provide function for proxy object
 */
func (this *provider) ProvideFunc(orginFun interface{}, fptr interface{}) error {

	// ---- create proxy ---
	this.provideProxyInjector(fptr, orginFun)

	proxyHandler := reflect.ValueOf(fptr).Elem().Interface()

	/**
	 * define the base fun name
	 */
	fn := funcName(orginFun)
	this.bindingFuns[fn] = proxyHandler

	return nil

}

// --- call and bind bean
func (this *provider) InjectBean(beanFun interface{}) error {

	fn := funcName(beanFun)

	this.beanFuns[fn] = beanFun

	return nil
}

func (this *provider) Invoke(handlers ...interface{}) error {
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
func (this *provider) String() string {
	return fmt.Sprintf("beanCtx{%s}", "update content ")
}

/**
 * create new beanCtx implement
 */
func createProvider() *provider {
	return &provider{make([]error, 0), nil, make(map[string]interface{}), make(map[reflect.Type]reflect.Type), make([]interface{}, 0), make(map[string]interface{})}
}

func (this *provider) provideProxyInjector(fptr interface{}, orgFun interface{}) {

	// --- define the target function element ----
	fn := reflect.ValueOf(fptr).Elem()

	refOrgFun := reflect.ValueOf(orgFun)

	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {

		var out []reflect.Value

		out = make([]reflect.Value, 0)

		callOut := refOrgFun.Call(in)

		// --- append out value ---

		for _, cOut := range callOut {
			out = append(out, cOut)
		}

		return out
	})

	fn.Set(v)
}
