package modules

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
	"log"
)

type Case1Module struct {
}

func (myself *Case1Module) Bind(ctx di.ModuleContext) {

	/**
	 * 注册提供使用的 struct bean
	 */
	ctx.GetRegister().RegBean(myself.provideCase1Obj1())

	/**
	 * 定义提供使用的 struct bean
	 */
	ctx.GetRegister().Invoke(myself.injectCase1)

}

/**
 * 定义注册的bean
 */
func (myself *Case1Module) provideCase1Obj1() *di.RegisterBean {

	var pro func() *mockobject.Case1MockObj1
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case1MockObj1{}
	rb.ProvideFun = &pro

	return rb
}

/**
 * 调用函数
 */
func (myself *Case1Module) injectCase1(obj1 *mockobject.Case1MockObj1) {
	log.SetPrefix("Case1Module.injectCase1: ")

	log.Print(" Start to invoke function  -> ")
	obj1.SayHello()
	log.Print(" End to invoke function  <- ")

}
