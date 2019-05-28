package services

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
)

type ServiceModule struct {
}

func (myself *ServiceModule) Bind(ctx di.ModuleContext) {

	/**
	 * 注册提供使用的 struct bean
	 */
	ctx.GetRegister().RegBean(myself.provideService())

	ctx.GetRegister().RegBean(myself.provideSayHello())

	/**
	 * 定义提供使用的 struct bean
	 */
	//ctx.GetRegister().Invoke(myself.injectCase1)

}

//
func (myself *ServiceModule) provideService() *di.RegisterBean {

	//var pro func() mockobject.SayHelloI
	//var pro func() *TestService
	rb := new(di.RegisterBean)
	rb.Bean = NewTestService()
	rb.ProvideFun = new(func() *TestService)
	return rb
}

func (myself *ServiceModule) provideSayHello() *di.RegisterBean {
	var case1 *mockobject.SayHelloCase1

	var pro = func() mockobject.SayHelloI {
		case1 = new(mockobject.SayHelloCase1)
		return case1
	}

	rb := new(di.RegisterBean)
	rb.Bean = case1
	rb.ProvideFun = &pro
	return rb
}
