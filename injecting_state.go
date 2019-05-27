package fire_di

import (
	"fmt"
	"github.com/1-bi/log-api"
	"go.uber.org/dig"
	"reflect"
)

type InjectingState struct {
	bean interface{}

	injectMethods map[reflect.Method]reflect.Value

	// this attribute is change inject candidate method
	injectCandidate map[reflect.Method]int

	// after set method define
	aftersetMethod reflect.Value

	logger logapi.Logger

	digContainer *dig.Container
}

func NewInjectingState(bean interface{}, injectMethods map[reflect.Method]reflect.Value, aftersetMethod reflect.Value) *InjectingState {

	state := new(InjectingState)
	state.bean = bean
	state.injectMethods = injectMethods
	state.aftersetMethod = aftersetMethod
	state.logger = logapi.GetLogger("fire-di")

	state.injectCandidate = make(map[reflect.Method]int)
	for m, _ := range injectMethods {
		state.injectCandidate[m] = 1
	}

	return state
}

func (myself *InjectingState) SetDigContainer(digContainer *dig.Container) {
	myself.digContainer = digContainer
}

func (myself *InjectingState) DoWork() {

	if myself.logger.IsDebugEnabled() {
		// --- get bean name -
		fmt.Println(reflect.ValueOf(myself.bean).String())

	}

}
