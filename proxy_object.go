package fire_di

import "reflect"

/**
 * defin all project object interface
 */
type proxyObject struct {
	ref interface{}

	methods map[string]reflect.Method
}
