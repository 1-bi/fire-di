package fire_di

import (
	"fmt"
	"reflect"
	"strings"
)

/**
 * defin all project object interface
 */
type proxyObject struct {
	ref     interface{}
	methods map[string]reflect.Method

	injectMethod map[string]reflect.Method
}

// apply proxy to self
func (myself *proxyObject) applyProxy(src interface{}) {

	objType := reflect.TypeOf(src)

	methodMap := make(map[string]reflect.Method, 0)
	var i int
	for i = 0; i < objType.NumMethod(); i++ {
		m := objType.Method(i)
		methodMap[m.Name] = m
	}
	myself.methods = methodMap

	// --- define bean ---
	myself.ref = src

}

type proxyInjectObject struct {
	ref interface{}
}

func (myself *proxyInjectObject) getProxyInvokeFun() interface{} {

	var tplFun func(ref interface{})

	targetFun := reflect.ValueOf(tplFun).Elem()

	v := reflect.MakeFunc(targetFun.Type(), func(args []reflect.Value) (results []reflect.Value) {

		fmt.Println("-hellomoo")
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
