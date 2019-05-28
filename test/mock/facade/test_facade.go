package facade

import (
	"github.com/1-bi/fire-di/test/mock/services"
	"github.com/1-bi/log-api"
	"log"
)

type TestFacade struct {
	testService *services.TestService

	logger logapi.Logger
}

/**
 * contruct instance
 */
func NewTestFacade() *TestFacade {
	taskFacade := new(TestFacade)
	return taskFacade
}

/**
 * inject base function
 * @Inject
 */
func (myself *TestFacade) InjectService(
	testService *services.TestService) {
	//myself.clientApi = clientApi

	myself.testService = testService

	myself.logger = logapi.GetLogger("fire-di.facade")

	//myself.servbusEvent = servbusEvent
	//myself.logger = logger

}

func (myself *TestFacade) TestFacadeMethod() {
	log.Println("Test facade method define ")
	myself.testService.TestMethod1()

}
