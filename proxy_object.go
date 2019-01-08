package fire_di

import "reflect"

/**
 * defin all project object interface
 */
type proxyObject struct {
	ref     interface{}
	methods map[string]reflect.Method
}

type proxyInjectObject struct {
	ref interface{}
}

func (myself *proxyInjectObject) getProxyInvokeFun() interface{} {
	return nil
}
