package fire_di

import (
	"fmt"
	"go.uber.org/dig"
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

	if err != nil {
		return injector, err
	}

	return injector, err
}

/**
 * define proxy message
 */
func (i *injector) scanProxyInject(proxies map[string]*proxyObject) error {

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
		fname := funcName(fn)

		err = i.container.Invoke(fn)

		if err != nil {
			fmt.Printf("Error during %q invoke: %v", fname, err)
			break
		}
	}

	return err
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
