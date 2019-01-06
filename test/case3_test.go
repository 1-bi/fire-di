package test

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
	tm "github.com/1-bi/fire-di/test/modules"
	"log"
	"testing"
)

/**
 * case 3 , add project
 */
func TestDI_module_case03(t *testing.T) {

	diConf := new(di.Configuration)

	/**
	 * custom  di config
	 */
	di.Config(diConf)

	module := tm.Case3Module{}

	// ----- register module pre defined ----
	bs := di.RegisterModules(&module)

	// ----- create injector for modules
	injector, err := di.CreateInjector(bs)

	// ----- create injector ----
	if err != nil {
		log.Fatal(err)
	} else {
		injector.Execute(bootstrapCase3)
	}

}

func bootstrapCase3(helper *mockobject.Case3MockObj1) {
	helper.SayHello()
}
