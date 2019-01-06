package modules

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/modules"
	"github.com/1-bi/fire-di/test/mockobject"
)

type Case3Module struct {
	Parent *modules.InjectSupportedModule
}

func (this *Case3Module) Bind(ctx di.ModuleContext) {
	// --- define function type , and assign provider method
	ctx.GetRegister().RegBean(this.provideCase3Obj2())
	ctx.GetRegister().RegBean(this.provideCase3Obj1())
	ctx.GetRegister().RegBean(this.provideCase3Obj3())
}

/**
 * define base the object method
 */
func (this *Case3Module) provideCase3Obj1() *di.RegisterBean {

	var pro func() *mockobject.Case3MockObj1
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case3MockObj1{}
	rb.ProvideFun = &pro

	return rb
}

func (this *Case3Module) provideCase3Obj2() *di.RegisterBean {

	var pro func(beanContextBinder *di.BeanCtxBinder) *mockobject.Case3MockObj2
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case3MockObj2{}
	rb.ProvideFun = &pro

	return rb
}

func (this *Case3Module) provideCase3Obj3() *di.RegisterBean {

	var pro func(beanContextBinder *di.BeanCtxBinder) *mockobject.Case3MockObj3
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case3MockObj3{}
	rb.ProvideFun = &pro

	return rb
}
