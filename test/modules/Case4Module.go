package modules

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
)

type Case4Module struct {
}

func (this *Case4Module) Bind(ctx di.ModuleContext) {
	// --- define function type , and assign provider method
	ctx.GetRegister().RegBean(this.provideCase4Obj2())

	ctx.GetRegister().RegFunc(this.provideCase3Obj1)

}

/**
 * define base the object method
 */
func (this *Case4Module) provideCase3Obj1(mock *mockobject.Case4MockObj2) *mockobject.Case4MockObj1 {

	return nil
}

func (this *Case4Module) provideCase4Obj2() *di.RegisterBean {

	var pro func(beanContextBinder *di.BeanCtxBinder) *mockobject.Case4MockObj2
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case4MockObj2{}
	rb.ProvideFun = &pro

	return rb
}
