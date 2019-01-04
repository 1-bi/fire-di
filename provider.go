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

func (this *provider) Provide(handlers ...interface{}) error {

	// ---- get all handles ----
	for _, handler := range handlers {

		fn := funcName(handler)

		this.bindingFuns[fn] = handler

		targetFun := handler

		// ---- create proxy ----
		var decoratedFunc, targetFunc reflect.Value
		targetFunc = reflect.ValueOf(handler)
		decoratedFunc = reflect.ValueOf(targetFun).Elem()

		v := reflect.MakeFunc(targetFunc.Type(),
			func(in []reflect.Value) (out []reflect.Value) {
				fmt.Println("before")
				out = targetFunc.Call(in)
				fmt.Println("after")
				return
			})
		decoratedFunc.Set(v)

		fmt.Println(v)

	}
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

// ======================= private method

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
