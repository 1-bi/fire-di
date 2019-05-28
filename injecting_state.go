package fire_di

import (
	"github.com/1-bi/log-api"
	"go.uber.org/dig"
	"log"
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
		//sb := logapi.NewStructBean()
		//sb.LogString("bean name ", reflect.ValueOf(myself.bean).String())
		//myself.logger.Debug("Inject bean .", sb)
	}

	var err error

	for m, methodRef := range myself.injectMethods {
		refTarFun := reflect.New(methodRef.Type())
		fn := refTarFun.Interface()

		// ---- set the value ---
		resultFun := FuncInterceptor(fn, func(in []reflect.Value) []reflect.Value {
			// call object
			result := methodRef.Call(in)

			// --- remove state
			// --- logic for after inject method
			delete(myself.injectCandidate, m)
			if len(myself.injectCandidate) == 0 {
				// --- fire the after method ---

				if myself.aftersetMethod.IsValid() && !myself.aftersetMethod.IsNil() {
					refValues := myself.callAftersetfun(myself.aftersetMethod)

					// run after set
					if myself.logger.IsDebugEnabled() {
						myself.logger.Debug("Show value after executing afterset function.", nil)
						for _, refVal := range refValues {
							log.Print(refVal)
						}
					}

				}
			}

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

func (myself *InjectingState) callAftersetfun(funAfter reflect.Value) []reflect.Value {
	if funAfter.Kind() != reflect.Invalid {
		return funAfter.Call(nil)
	}
	return nil
}
