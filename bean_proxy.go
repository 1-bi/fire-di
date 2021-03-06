package fire_di

import (
	"reflect"
	"strings"
)

/**
 * defin all project object interface
 */
type BeanProxy struct {
	ref     interface{}
	methods map[string]reflect.Method

	injectMethods map[reflect.Method]reflect.Value
	// after set method define
	aftersetMethod reflect.Value
}

// apply proxy to self
func (myself *BeanProxy) applyProxy(src interface{}) {

	objType := reflect.TypeOf(src)
	objRev := reflect.ValueOf(src)

	methodMap := make(map[string]reflect.Method, 0)
	injectMap := make(map[reflect.Method]reflect.Value, 0)

	var i int
	for i = 0; i < objType.NumMethod(); i++ {
		m := objType.Method(i)
		methodMap[m.Name] = m

		var matchPrefix bool

		// --- check the use define prefix method prefix ---
		for _, prefix := range runnimeConf.injectMethodPrefix {

			if strings.HasPrefix(m.Name, prefix) {
				matchPrefix = true
				break
			}
		}

		if matchPrefix {
			// ---- invoke method bean ---

			injectMap[m] = objRev.Method(i)

		}

		if strings.Compare(m.Name, "Afterset") == 0 {
			myself.aftersetMethod = objRev.Method(i)

		}

	}
	myself.methods = methodMap
	myself.injectMethods = injectMap

	// --- define bean ---
	myself.ref = src

}

func (myself *BeanProxy) CreateInjectingState() *InjectingState {
	return NewInjectingState(myself.ref, myself.injectMethods, myself.aftersetMethod)
}

type proxyInjectObject struct {
	ref interface{}
}

func (myself *proxyInjectObject) getProxyInvokeFun() interface{} {

	var tplFun func(ref interface{})

	targetFun := reflect.ValueOf(tplFun).Elem()

	v := reflect.MakeFunc(targetFun.Type(), func(args []reflect.Value) (results []reflect.Value) {

		if len(args) == 0 {
			return nil
		}

		var ret reflect.Value

		switch args[0].Kind() {
		case reflect.Int:
			n := 0
			for _, a := range args {
				n += int(a.Int())
			}

			ret = reflect.ValueOf(n)
		case reflect.String:
			ss := make([]string, 0, len(args))
			for _, s := range args {
				ss = append(ss, s.String())
			}

			ret = reflect.ValueOf(strings.Join(ss, ""))
		}

		results = append(results, ret)
		return
	})

	targetFun.Set(v)

	return targetFun.Interface()
}
