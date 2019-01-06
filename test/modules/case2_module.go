package modules

import (
	"errors"
	"fmt"
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
)

type Case2Module struct {
}

func (this *Case2Module) Bind(ctx di.ModuleContext) {

	ctx.GetRegister().Provide(this.provideCase2Obj1)
	ctx.GetRegister().Provide(this.provideCase2Obj2)
	//ctx.GetRegister().Provide( this.provideCase2Obj3)

	ctx.GetRegister().InjectBean(this.injectCase2)

}

func (this *Case2Module) provideCase2Obj1() (*mockobject.Case2MockObj1, error) {
	case1 := mockobject.Case2MockObj1{}

	return &case1, errors.New("hello message.")
}

func (this *Case2Module) provideCase2Obj2() *mockobject.Case2MockObj2 {
	case1 := mockobject.Case2MockObj2{}
	return &case1
}

func (this *Case2Module) provideCase2Obj3() *mockobject.Case2MockObj3 {
	case1 := mockobject.Case2MockObj3{}
	return &case1
}

func (this *Case2Module) injectCase2(obj1 *mockobject.Case2MockObj1, obj2 *mockobject.Case2MockObj2, obj3 *mockobject.Case2MockObj3) {
	fmt.Println("hello message")
}
