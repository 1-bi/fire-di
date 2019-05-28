package services

import (
	"github.com/1-bi/fire-di/test/mockobject"
	"github.com/1-bi/log-api"
	"log"
)

/**
 * contruct instance
 */
func NewTestService() *TestService {
	taskServ := new(TestService)
	return taskServ
}

type TestService struct {
	logger logapi.Logger

	sayHello mockobject.SayHelloI
}

// InjectSayHello inject say hello function
func (myself *TestService) InjectSayHello(sayHello mockobject.SayHelloI) {
	myself.sayHello = sayHello
}

//  CreateTask create task for input
func (myself *TestService) TestMethod1() {

	log.Println("call test service method1")

	myself.sayHello.SayHello("tester ")

}
