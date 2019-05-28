package facade

import (
	di "github.com/1-bi/fire-di"
)

type FacadeModule struct {
}

func (myself *FacadeModule) Bind(ctx di.ModuleContext) {

	/**
	 * 注册提供使用的 struct bean
	 */
	ctx.GetRegister().RegBean(myself.provideFacade())

	/**
	 * 定义提供使用的 struct bean
	 */
	//ctx.GetRegister().Invoke(myself.injectCase1)

}

/**
 * 定义注册的bean
 */
func (myself *FacadeModule) provideFacade() *di.RegisterBean {

	//var pro func() mockobject.SayHelloI
	var pro func() *TestFacade
	rb := new(di.RegisterBean)
	rb.Bean = new(TestFacade)
	rb.ProvideFun = &pro

	return rb
}
