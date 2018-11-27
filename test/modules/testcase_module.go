package modules

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
)

type Case1Module struct {
}

func (this *Case1Module) Bind(ctx di.ModuleContext) {

	ctx.GetProvider().Provide(this.provideCase1Helper)
}

func (this *Case1Module) provideCase1Helper() *mockobject.Case1Helper {
	case1 := mockobject.Case1Helper{}
	return &case1
}
