package fire_di

import (
	"fmt"
	"go.uber.org/dig"
	"reflect"
	"strings"
)

type injector struct {
	container   *dig.Container
	invokedFuns []interface{}
}

// ---- call and create injector ---
func createInjector(bs providerstore) (*injector, error) {
	var err error
	// ---- create injector ----
	injector := &injector{dig.New(), bs.modContext.Provider.invokedFuns}

	// ---- create application ctx bean first -----
	err = injector.container.Provide(createBeanCtxBinderAndApplicationCtx)
	if err != nil {
		return injector, err
	}

	// ---- scan all object first ----

	//  ----- scan and add method to container ---
	for _, handler := range bs.modContext.GetRegister().bindingFuns {

		err = injector.container.Provide(handler)
		if err != nil {
			break
		}
	}

	if err != nil {
		return injector, err
	}

	// ---- handle and pass bean to application ctx
	err = injector.container.Invoke(initApplicationContextFromBeanContext)

	// ---- call register function ----
	for _, iFunc := range bs.modContext.GetRegister().beanFuns {
		err = injector.container.Invoke(iFunc)
		if err != nil {
			break
		}
	}

	// ---- get all injector object dependency mapping ---

	err = injector.proxyBeanInvokedFunDefined(bs.modContext.GetRegister().GetProxyBeans())

	if err != nil {
		return injector, err
	}

	return injector, err
}

func (myself *injector) proxyBeanInvokedFunDefined(proxyBeans []*BeanProxy) error {

	var err error

	for _, proxyBean := range proxyBeans {

		// provide proxy bean by function define
		beanInjectState := proxyBean.CreateInjectingState()
		// inject container
		beanInjectState.SetDigContainer(myself.container)

		err = beanInjectState.DoWork()

		if err != nil {
			break
		}

		/*
			dependencyStateArray := make([]*dependencyState, 0)

			for _, dependency := range proxyBean.dependentStructs {

				dependencyStateArray = append(dependencyStateArray, newDependencyState(dependency))
			}

			if len(dependencyStateArray) > 0 {

				myself.setProxyBeanInjectFun(proxyBean, dependencyStateArray)

			} else {
				//  call after method directory
				myself.callAftersetfun(proxyBean)
			}*/

	}
	return err
}

func (myself *injector) setProxyBeanInjectFun(proxyBean *BeanProxy, depenMethods []*dependencyState) {

	for _, methodRef := range proxyBean.injectMethods {

		refTarFun := reflect.New(methodRef.Type())
		fn := refTarFun.Interface()

		// ---- set the value ---
		resultFun := FuncInterceptor(fn, func(in []reflect.Value) []reflect.Value {
			// call object
			result := methodRef.Call(in)

			// --- defined depenMethods status ---
			for _, state := range depenMethods {

				for _, inCls := range in {

					if state.stateInjected == 0 && inCls.Type().String() == state.dependencyType {
						state.updateState(1)
						break
					}
				}

			}

			// --- check all dependency class is load ---
			var allDepLoaded bool
			allDepLoaded = true
			for _, state := range depenMethods {
				// check dependency load or not
				if state.stateInjected == 0 {
					allDepLoaded = false
					break
				}
			}

			if allDepLoaded {
				// --- fire after event ---
				myself.callAftersetfun(proxyBean)
			}

			return result
		})

		myself.container.Invoke(resultFun.Interface())
	}

}

func (myself *injector) callAftersetfun(proxyBean *BeanProxy) {
	funAfter := proxyBean.aftersetMethod

	if funAfter.Kind() != reflect.Invalid {
		funAfter.Call(nil)
	}
}

/**
 * define proxy message
 */
func (i *injector) scanProxyInject(proxies map[string]*BeanProxy) error {

	for proxyName, proxyRef := range proxies {

		fmt.Println(proxyName)
		fmt.Println(proxyRef)

	}

	return nil
}

/**
 * define bootstrap handle event
 */
func (i *injector) Execute(funcs ...interface{}) error {

	var err error

	// --- call invoke function
	for _, fn := range i.invokedFuns {
		fname := funcName(fn)
		err = i.container.Invoke(fn)
		if err != nil {
			fmt.Printf("Error during %q invoke: %v", fname, err)
			break
		}
	}

	for _, fn := range funcs {

		// ---- create proxy function ---

		fname := funcName(fn)

		err = i.container.Invoke(fn)

		if err != nil {
			fmt.Printf("Error during %q invoke: %v", fname, err)
			break
		}
	}

	return err
}

/**
 * not use object
 */
func getFunProxy(orgFn interface{}) interface{} {

	fprtTyp := reflect.TypeOf(orgFn)

	targetFun := reflect.New(fprtTyp).Elem()

	v := reflect.MakeFunc(targetFun.Type(), func(args []reflect.Value) (results []reflect.Value) {

		fmt.Println("-hellomoo")
		if len(args) == 0 {
			return nil
		}

		var ret reflect.Value

		switch args[0].Kind() {
		case reflect.Int:
			n := 0
			for _, a := range args {
				n += int(a.Int())
			}

			ret = reflect.ValueOf(n)
		case reflect.String:
			ss := make([]string, 0, len(args))
			for _, s := range args {
				ss = append(ss, s.String())
			}

			ret = reflect.ValueOf(strings.Join(ss, ""))
		}

		results = append(results, ret)
		return
	})

	targetFun.Set(v)

	return targetFun.Interface()

}

// ================== private function area ====================
func createBeanCtxBinderAndApplicationCtx() (beanContextBinder *BeanCtxBinder, appCtx *AppCtx) {
	// --- init bean context binder object ----
	beanContextBinder = createBeanCtxBinder()
	appCtx = createAppCtx()
	return beanContextBinder, appCtx

}

/**
 * init base call application ctx binder
 */
func initApplicationContextFromBeanContext(beanCtx *BeanCtxBinder, appCtx *AppCtx) {

	for bindingKey, bindingRefObj := range beanCtx.bindings {

		resolvedBinding, err := bindingRefObj.resolvedBinding(nil, nil)
		if err != nil {

		}
		appCtx.bindings[bindingKey] = resolvedBinding
	}

}
