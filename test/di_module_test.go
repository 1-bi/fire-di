package test

import (
	"fmt"
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/modules"
	"testing"
)

/**
 * sample case 01 for generate id
 */
func TestDI_module_case01(t *testing.T) {

	// ----- register module pre defined ----
	bs := di.RegisterModules(&modules.TestCaeeModule{})

	// ----- create injector for modules
	injector := di.CreateInjector(bs)

	// ----- create injector ----
	fmt.Println(injector)

}
