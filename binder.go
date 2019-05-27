package fire_di

import (
	"fmt"
	"reflect"
)

// BeanCtxBinder  defined appp store
type BeanCtxBinder struct {
	bindingErrors []error
	bindings      map[bindingKey]binding

	aliasNameMapping map[string]reflect.Type
	bindBeans        map[reflect.Type]*BeanProxy
}

// Bind ... bind all interface
func (myself *BeanCtxBinder) Bind(froms ...interface{}) Builder {
	// ---- create builder ----
	return myself.bind(newBindingKey, froms)
}

/**
 *  define alias type name
 */
func (myself *BeanCtxBinder) alias(aliasType string, dstType reflect.Type) {
	myself.aliasNameMapping[aliasType] = dstType
}

func (myself *BeanCtxBinder) bindProxyInst(proxyBean *BeanProxy, refType reflect.Type) {
	myself.bindBeans[refType] = proxyBean
}

// -------------------- bind context -------------
/**
 * bind function inteface with reflect type
 */
func (myself *BeanCtxBinder) bind(newBindingKeyFunc func(reflect.Type) bindingKey, from []interface{}) InterfaceBuilder {

	lenFrom := len(from)
	if lenFrom == 0 {
		myself.addBindingError(errNil)
		return newNoOpBuilder()
	}

	bindingKeys := make([]bindingKey, lenFrom)

	for i := 0; i < lenFrom; i++ {

		fromReflectType := reflect.TypeOf(from[i])

		if fromReflectType == nil {
			myself.addBindingError(errNil)
			return newNoOpBuilder()
		}

		fromOrgType := fromReflectType.Elem()

		// --- customer array ---
		bindingKeys[i] = newBindingKeyFunc(fromOrgType)
	}

	return newBuilder(myself, bindingKeys)
}

func (myself *BeanCtxBinder) addBindingError(err error) {
	myself.bindingErrors = append(myself.bindingErrors, err)
}

// BindType bind type in builder
func (myself *BeanCtxBinder) BindType(froms ...interface{}) Builder {

	// ---- create builder ----
	return myself.bind(newBindingKey, froms)
}

/**
 * defined method binding
 * implement function for inject api
 */
func (myself *BeanCtxBinder) String() string {
	return fmt.Sprintf("beanCtx{%s}", "update content ")
}

/**
 * create new beanCtx implement
 */
func createBeanCtxBinder() *BeanCtxBinder {
	return &BeanCtxBinder{make([]error, 0), make(map[bindingKey]binding),
		make(map[string]reflect.Type, 0), make(map[reflect.Type]*BeanProxy, 0)}
}

/**
 * set binding
 */
func (myself *BeanCtxBinder) setBinding(bindingKey bindingKey, binding binding) {
	foundBinding, ok := myself.bindings[bindingKey]
	if ok {
		myself.addBindingError(errAlreadyBound.withTag("bindingKey", bindingKey).withTag("foundBinding", foundBinding))
		return
	}
	myself.bindings[bindingKey] = binding
}
