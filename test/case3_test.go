package test

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/modules"
	"github.com/1-bi/fire-di/test/mockobject"
	tm "github.com/1-bi/fire-di/test/modules"
	"log"
	"testing"
)

/**
 * sample case 01 for generate id
 */
func TestDI_module_case03(t *testing.T) {

	annoMod := new(modules.InjectSupportedModule)
	module := tm.AnnotationSupportedModule{annoMod}

	// ----- register module pre defined ----
	bs := di.RegisterModules(&module)

	// ----- create injector for modules
	injector, err := di.CreateInjector(bs)

	// ----- create injector ----
	if err != nil {
		log.Fatal(err)
	} else {
		injector.Execute(bootstrapCase2)
	}

}

func bootstrapCase3(helper *mockobject.Case2MockObj1) {
	helper.SayHello()
}
