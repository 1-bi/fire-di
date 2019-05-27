package fixture

import (
	"fmt"
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/fire-di/test/mockobject"
	"github.com/1-bi/fire-di/test/modules"
	"github.com/1-bi/log-api"
	"github.com/smartystreets/gunit"
	"log"
)

type InterfaceFixTure struct {
	*gunit.Fixture

	logger logapi.Logger
}

// Setup
func (myself *InterfaceFixTure) Setup() {

	myself.logger = logapi.GetLogger("main")

	/*	// ----- register module pre defined ----

		// ----- create injector for modules
		injector, err := di.CreateInjector(bs)

		// ----- create injector ----
		if err != nil {
			log.Fatal(err)
		} else {
			injector.Execute(myself.invokerFun)
		}
	*/
	fmt.Println("----")
}

func (myself *InterfaceFixTure) Teardown() {

}

func (myself *InterfaceFixTure) TestCase1() {

	diConf := new(di.Configuration)

	/**
	 * custom  di config
	 */
	di.Config(diConf)

	mod := new(modules.InterfaceModule)
	bs := di.RegisterModules(mod)

	// ----- create injector for modules
	injector, err := di.CreateInjector(bs)

	// ----- create injector ----
	if err != nil {
		log.Fatal(err)
	} else {
		injector.Execute(myself.invokerFun)
	}
}

// --- test intereface mesage
func (myself *InterfaceFixTure) invokerFun(sayHello mockobject.SayHelloI) {
	myself.logger.Info("call say hello function ", nil)

	sayHello.SayHello(" myname is who. ")

}
