package fire_di

import (
	"reflect"
)

var (
	errorReflectType = reflect.TypeOf((*error)(nil)).Elem()
)

/**
 * create empty builder object
 */
type noOpBuilder struct{}

func newNoOpBuilder() InterfaceBuilder {
	return &noOpBuilder{}
}

func (n *noOpBuilder) ToProxyInst(singleton reflect.Value) {}

/**
 * create base builder object
 */
type baseBuilder struct {
	beanCtx     *BeanCtxBinder
	bindingKeys []bindingKey
}

func newBuilder(binder *BeanCtxBinder, bindingKeys []bindingKey) InterfaceBuilder {
	return &baseBuilder{binder, bindingKeys}
}

/**
 * bind base proxy
 */
func (myself *baseBuilder) getProxy(ref interface{}) *InjectObjInfoProxy {
	proxyObj := new(InjectObjInfoProxy)
	proxyObj.ref = ref
	objType := reflect.TypeOf(ref)

	methodMap := make(map[string]reflect.Method, 0)
	var i int
	for i = 0; i < objType.NumMethod(); i++ {
		m := objType.Method(i)
		methodMap[m.Name] = m
	}
	proxyObj.methods = methodMap

	return proxyObj
}

/**
 * create proxy instance
 */
func (myself *baseBuilder) ToProxyInst(singletonVal reflect.Value) {
	// use instance for singleton
	//myself.to(singleton, verifyInterfaceReflectType, newSingletonBinding)

	// create proxy instance
	proxyObj := myself.getProxy(singletonVal.Interface())

	// --- check binding object
	myself.beanCtx.bindProxyInst(proxyObj, singletonVal.Type())

}

func (myself *baseBuilder) to(object reflect.Value, verifyFunc func(reflect.Type, reflect.Type) error, newBindingFunc func(interface{}) binding) {
	objectReflectType := object.Type()

	for _, bindingKey := range myself.bindingKeys {
		if err := verifyFunc(bindingKey.reflectType(), objectReflectType); err != nil {
			myself.beanCtx.addBindingError(err)
			return
		}
	}

	binding := newBindingFunc(object)

	// ---- output binding value ---
	for _, bindingKey := range myself.bindingKeys {
		myself.setBinding(bindingKey, binding)
	}

}

func (myself *baseBuilder) setBinding(bindingKey bindingKey, binding binding) {

	myself.beanCtx.setBinding(bindingKey, binding)

}

func verifyInterfaceReflectType(bindingKeyReflectType reflect.Type, bindingReflectInstanceType reflect.Type) error {
	// --- check the interface type match ---
	if !bindingReflectInstanceType.Implements(bindingKeyReflectType) {
		return errImplementNotSuit.withTag("implementationReflectType", bindingReflectInstanceType)
	}
	return nil
}

/**
 * mark value ---
 */
func verifyBindingReflectType(bindingKeyReflectType reflect.Type, bindingReflectType reflect.Type) error {

	if !isSupportedBindingKeyReflectType(bindingKeyReflectType) {
		return errNotSupportedYet.withTag("bindingKeyReflectType", bindingReflectType)
	}
	if isInterfacePtr(bindingKeyReflectType) {
		bindingKeyReflectType = bindingKeyReflectType.Elem()
	}
	if !bindingReflectType.AssignableTo(bindingKeyReflectType) {
		return errNotAssignable.withTag("bindingKeyReflectType", bindingKeyReflectType).withTag("bindingReflectType", bindingReflectType)
	}
	return nil
}

func verifyConstructorReturnValues(bindingKeyReflectType reflect.Type, constructorReflectType reflect.Type) error {
	if constructorReflectType.NumOut() != 2 {
		return errConstructorReturnValuesInvalid.withTag("constructorReflectType", constructorReflectType)
	}
	if err := verifyBindingReflectType(bindingKeyReflectType, constructorReflectType.Out(0)); err != nil {
		return err
	}
	if !constructorReflectType.Out(1).AssignableTo(errorReflectType) {
		return errConstructorReturnValuesInvalid.withTag("constructorReflectType", constructorReflectType)
	}
	return nil
}
