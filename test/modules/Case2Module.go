package modules

import (
	"fmt"
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
)

type Case2Module struct {
}

func (myself *Case2Module) Bind(ctx di.ModuleContext) {

	ctx.GetRegister().RegBean(myself.provideCase2Obj1())
	ctx.GetRegister().RegBean(myself.provideCase2Obj2())
	ctx.GetRegister().RegBean(myself.provideCase2Obj3())

	ctx.GetRegister().InjectBean(myself.injectCase2)

}

func (myself *Case2Module) provideCase2Obj1() *di.RegisterBean {

	var pro func() *mockobject.Case2MockObj1
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case2MockObj1{}
	rb.ProvideFun = &pro

	return rb
}

func (myself *Case2Module) provideCase2Obj2() *di.RegisterBean {

	var pro func() mockobject.MockInterface
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case2MockObj2{}
	rb.ProvideFun = &pro

	return rb
}

func (myself *Case2Module) provideCase2Obj3() *di.RegisterBean {

	var pro func() *mockobject.Case2MockObj3
	rb := new(di.RegisterBean)
	rb.Bean = &mockobject.Case2MockObj3{}
	rb.ProvideFun = &pro

	return rb
}

func (myself *Case2Module) injectCase2(obj1 *mockobject.Case2MockObj1, obj2 mockobject.MockInterface, obj3 *mockobject.Case2MockObj3) {
	fmt.Println("hello message")

	obj2.TestMock()
}
