package fire_di

import (
	"fmt"
	"gitlab.com/1-bi/log-api/loggercom"
	"reflect"
	"strings"
)

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
 * define register bean
 */
type RegisterBean struct {
	// define base bean
	Bean interface{}

	ProvideFun interface{}
}

/**
 * fix new bug
 */
// Deprecated:  this method would be remove in next version
func (this *register) Provide(handlers ...interface{}) error {

	var err error

	// ---- get all handles ----
	for _, handler := range handlers {

		err = this.privateFun(handler)
		if err != nil {
			return err
		}

	}
	return nil
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

		if !strings.HasPrefix(methodName, "Inject") {
			continue
		}

		// ---- invoke method bean ---
		this.Invoke(m.Func.Interface())
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

/**
 * add provide function for proxy object
 * fire the event when add the provide in method ---
 */
func (this *register) privateFun(orginFun interface{}) error {

	orgVal := reflect.ValueOf(orginFun)
	fptrRefVal := reflect.New(orgVal.Type())

	// ---- create proxy ---
	this.provideProxyInjector(fptrRefVal, orginFun)

	proxyHandler := fptrRefVal.Elem().Interface()

	/**
	 * define the base fun name
	 */
	fn := funcName(orginFun)
	this.bindingFuns[fn] = proxyHandler

	return nil

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

func (this *register) provideProxyInjector(fptr reflect.Value, orgFun interface{}) {

	// --- define the target function element ----
	fn := fptr.Elem()

	refOrgFun := reflect.ValueOf(orgFun)

	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {

		// --- check the bean context ---
		/*
			var bcb *BeanCtxBinder
			for _, input := range in {

				if reflect.TypeOf( (*BeanCtxBinder)(nil) ) == input.Type() {
					bcb = input.Interface().(*BeanCtxBinder)
				}
			}

			if bcb == nil {
				// --- throw error ----
			}
		*/
		var out []reflect.Value
		out = make([]reflect.Value, 0)
		callOut := refOrgFun.Call(in)

		// --- append out value ---
		for _, cOut := range callOut {

			// --- set the alias name ---
			/*
				if reflect.Ptr == cOut.Kind() {
					bcb.alias( cOut.Elem().Type().String() , cOut.Type() )
				}
				builder := bcb.Bind(  cOut.Type() )
				builder.ToProxyInst( cOut )
			*/

			out = append(out, cOut)
			// ---- chck error -----

		}
		//fmt.Println( out  )

		return out
	})

	fn.Set(v)
}
