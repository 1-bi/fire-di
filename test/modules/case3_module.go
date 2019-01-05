package modules

import (
	"fmt"
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/modules"
	"github.com/1-bi/fire-di/test/mockobject"
	"reflect"
)

type Case3Module struct {
	Parent *modules.InjectSupportedModule
}

func (this *Case3Module) Bind(ctx di.ModuleContext) {

	var proCase3Obj1 func() *mockobject.Case3MockObj1

	/*
		makeProxy := func(fptr interface{}, funRef interface{}) {

			fn := reflect.ValueOf(fptr).Elem()
			v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {

				var  out []reflect.Value

				out = make([]reflect.Value,0)

				funElem := reflect.ValueOf( funRef )

				callOut := funElem.Call(nil)

				out = append(out , callOut[0])


				return out
			})

			fn.Set(v)


		}

		// --- inter face mapping ---
		makeProxy(&proCase3Obj1 , this.provideCase3Obj1)
	*/

	// --- call object ---
	//ctx.GetProvider().Provide(proCase3Obj1)
	ctx.GetProvider().ProvideFunc(this.provideCase3Obj1, &proCase3Obj1)
	fmt.Println(proCase3Obj1)

	fmt.Println(&proCase3Obj1)
	fmt.Println(reflect.ValueOf(&proCase3Obj1).Elem())

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
