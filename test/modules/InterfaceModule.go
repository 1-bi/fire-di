package modules

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
	"log"
)

type InterfaceModule struct {
}

func (myself *InterfaceModule) Bind(ctx di.ModuleContext) {

	/**
	 * 注册提供使用的 struct bean
	 */
	ctx.GetRegister().RegBean(myself.provideInterface_Case1())

	/**
	 * 定义提供使用的 struct bean
	 */
	//ctx.GetRegister().Invoke(myself.injectCase1)

}

/**
 * 定义注册的bean
 */
func (myself *InterfaceModule) provideInterface_Case1() *di.RegisterBean {

	//var pro func() mockobject.SayHelloI
	var pro func() *mockobject.SayHelloCase1
	rb := new(di.RegisterBean)
	rb.Bean = new(mockobject.SayHelloCase1)
	rb.ProvideFun = &pro

	return rb
}

/**
 * 调用函数
 */
func (myself *InterfaceModule) injectCase1(obj1 *mockobject.Case1MockObj1) {
	log.SetPrefix("InterfaceModule.injectCase1: ")

	log.Print(" Start to invoke function  -> ")
	obj1.SayHello()
	log.Print(" End to invoke function  <- ")

}
