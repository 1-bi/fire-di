package modules

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/modules"
	"github.com/1-bi/fire-di/test/mockobject"
)

type AnnotationSupportedModule struct {
	parent *modules.InjectSupportedModule
}

func (this *AnnotationSupportedModule) Bind(ctx di.ModuleContext) {

	// --- call parent method ---
	this.parent.Bind(ctx)

}

func (this *AnnotationSupportedModule) provideCase2Obj1() *mockobject.Case2MockObj1 {
	case1 := mockobject.Case2MockObj1{}

	// --- register bean for inject ----
	this.parent.RegisterBean(&case1)
	return &case1
}

func (this *AnnotationSupportedModule) provideCase2Obj2() *mockobject.Case2MockObj2 {
	case1 := mockobject.Case2MockObj2{}
	return &case1
}

func (this *AnnotationSupportedModule) provideCase2Obj3() *mockobject.Case2MockObj3 {
	case1 := mockobject.Case2MockObj3{}
	return &case1
}
