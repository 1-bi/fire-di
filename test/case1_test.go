package test

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
	"github.com/1-bi/fire-di/test/modules"
	"log"
	"testing"
)

/**
 * sample case 01 for generate id
 */
func TestDI_module_case01(t *testing.T) {

	// ----- register module pre defined ----
	bs := di.RegisterModules(&modules.Case1Module{})

	// ----- create injector for modules
	injector, err := di.CreateInjector(bs)

	// ----- create injector ----
	if err != nil {
		log.Fatal(err)
	} else {
		injector.Execute(bootstrapCase1)
	}
}

func bootstrapCase1(helper *mockobject.Case1Helper) {
	helper.SayHello()
}
