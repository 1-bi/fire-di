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

	injectMethods map[string]reflect.Method

	aftersetMethod reflect.Method

	dependentStructs []string
}

// apply proxy to self
func (myself *proxyObject) applyProxy(src interface{}) {

	objType := reflect.TypeOf(src)

	methodMap := make(map[string]reflect.Method, 0)
	injectMap := make(map[string]reflect.Method, 0)

	var i int
	for i = 0; i < objType.NumMethod(); i++ {
		m := objType.Method(i)
		methodMap[m.Name] = m

		// --- check method with prefix "Inject" ---
		if strings.HasPrefix(m.Name, "Inject") {

			// found object dependency
			myself.foundDependencyCls(m)

			fmt.Println(myself.dependentStructs)

			injectMap[m.Name] = m
		}

		if strings.Compare(m.Name, "Afterset") == 0 {
			myself.aftersetMethod = m
		}

	}
	myself.methods = methodMap
	myself.injectMethods = injectMap

	// --- define bean ---
	myself.ref = src

}

func (myself *proxyObject) foundDependencyCls(injectMethod reflect.Method) {

	methodTyp := injectMethod.Type
	var i int
	var existedDependency bool
	for i = 0; i < methodTyp.NumIn(); i++ {
		objectTyp := methodTyp.In(i)
		if objectTyp.Kind() == reflect.Ptr {

			// --- check the dependency ---
			existedDependency = false
			for _, v := range myself.dependentStructs {

				if v == objectTyp.String() {
					existedDependency = true
				}

			}

			if !existedDependency {
				myself.dependentStructs = append(myself.dependentStructs, objectTyp.String())
			}

		}
	}

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
