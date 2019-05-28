package test

import (
	"fmt"
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mock/facade"
	"github.com/1-bi/fire-di/test/mock/services"
	"hash/fnv"
	"log"
	"testing"
)

/**
 * case 3 , add project
 */
func TestDI_module_Inject(t *testing.T) {

	diConf := new(di.Configuration)

	/**
	 * custom  di config
	 */
	di.Config(diConf)

	// ----- register module pre defined ----
	bs := di.RegisterModules(
		&facade.FacadeModule{},
		&services.ServiceModule{},
	)

	// ----- create injector for modules
	injector, err := di.CreateInjector(bs)

	// ----- create injector ----
	if err != nil {
		log.Fatal(err)
	} else {
		injector.Execute(injectRun)
	}

}

func injectRun(testFacade *facade.TestFacade) {
	testFacade.TestFacadeMethod()
}

func Test_string(t *testing.T) {
	s := "sha1 tringfyy"

	h := fnv.New64()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Println(h.Sum64())
}
