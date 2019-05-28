package facade

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
)

type FacadeModule struct {
}

func (myself *FacadeModule) Bind(ctx di.ModuleContext) {

	/**
	 * 注册提供使用的 struct bean
	 */
	ctx.GetRegister().RegBean(myself.provideFacade())

	ctx.GetRegister().RegBean(myself.provideSayGoodbye())

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

func (myself *FacadeModule) provideSayGoodbye() *di.RegisterBean {

	//var pro func() mockobject.SayHelloI
	rb := new(di.RegisterBean)
	rb.Bean = new(mockobject.SayGoodbyeCase1)
	rb.ProvideFun = new(func() mockobject.GoodbyeI)
	return rb
}
