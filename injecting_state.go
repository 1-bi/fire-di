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

func (myself *InjectingState) DoWork() error {

	if myself.logger.IsDebugEnabled() {
		// --- get bean name -
		fmt.Println(reflect.ValueOf(myself.bean).String())
	}

	var err error

	for _, methodRef := range myself.injectMethods {
		refTarFun := reflect.New(methodRef.Type())
		fn := refTarFun.Interface()

		// ---- set the value ---
		resultFun := FuncInterceptor(fn, func(in []reflect.Value) []reflect.Value {
			// call object
			result := methodRef.Call(in)

			// --- define error object

			return result
		})

		// define function interface and mapping
		err = myself.digContainer.Invoke(resultFun.Interface())
		if err != nil {
			break
		}
	}
	return err

}
