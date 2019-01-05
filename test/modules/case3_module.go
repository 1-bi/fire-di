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
	ctx.GetProvider().Provide(this.provideCase3Obj1)

	//ctx.GetProvider().Provide( proCase3Obj1 )

	//ctx.GetProvider().Provide( reflect.ValueOf( proCase3Obj1 ).Interface() )
	//ctx.GetProvider().Provide( reflect.ValueOf( &proCase3Obj1 ).Elem().Interface() )

	//ctx.GetProvider().Provide(provideCase3Obj1)

	// --- call parent method ---
	//this.parent.Bind(ctx)

}

func (this *Case3Module) provideCase3Obj1() *mockobject.Case3MockObj1 {
	case1 := mockobject.Case3MockObj1{}

	// --- register bean for inject ----
	//this.parent.RegisterBean(&case1)
	return &case1
}

func (this *Case3Module) provideCase3Obj2() *mockobject.Case3MockObj2 {
	case1 := mockobject.Case3MockObj2{}
	return &case1
}

func (this *Case3Module) provideCase3Obj3() *mockobject.Case3MockObj3 {
	case1 := mockobject.Case3MockObj3{}
	return &case1
}
