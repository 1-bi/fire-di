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
func TestDI_module_case02(t *testing.T) {

	// ----- register module pre defined ----
	bs := di.RegisterModules(&modules.Case2Module{})

	// ----- create injector for modules
	injector, err := di.CreateInjector(bs)

	// ----- create injector ----
	if err != nil {
		log.Fatal(err)
	} else {
		injector.Execute(bootstrapCase2)
	}

}

func bootstrapCase2(helper *mockobject.Case2MockObj1) {
	helper.SayHello()
}
