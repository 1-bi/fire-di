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
func createInjector(bs providerstore) *injector {
	// ---- create injector ----
	injector := &injector{dig.New(), bs.modContext.Provider.invokedFuns}

	// ---- create application ctx bean first -----
	injector.container.Provide(createBeanCtxBinderAndApplicationCtx)

	//  ----- scan and add method to container ---

	for _, handler := range bs.modContext.Provider.bindingFuns {

		injector.container.Provide(handler)

	}

	// ---- call register function ----
	for _, iFunc := range bs.modContext.Provider.beanFuns {
		injector.container.Invoke(iFunc)
	}

	// ---- handle and pass bean to application ctx
	injector.container.Invoke(passApplicationContextFromBeanContext)

	return injector
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

func passApplicationContextFromBeanContext(beanCtx *BeanCtxBinder, appCtx *AppCtx) {

	for bindingKey, bindingRefObj := range beanCtx.bindings {

		resolvedBinding, err := bindingRefObj.resolvedBinding(nil, nil)
		if err != nil {

		}
		appCtx.bindings[bindingKey] = resolvedBinding
	}

}
