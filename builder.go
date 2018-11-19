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

func (n *noOpBuilder) ToSingleton(singleton interface{}) {}

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

func (this *baseBuilder) ToSingleton(singleton interface{}) {
	// use instance for singleton
	this.to(singleton, verifyInterfaceReflectType, newSingletonBinding)

}

func (this *baseBuilder) to(object interface{}, verifyFunc func(reflect.Type, reflect.Type) error, newBindingFunc func(interface{}) binding) {
	objectReflectType := reflect.TypeOf(object)

	for _, bindingKey := range this.bindingKeys {
		if err := verifyFunc(bindingKey.reflectType(), objectReflectType); err != nil {
			this.beanCtx.addBindingError(err)
			return
		}
	}

	binding := newBindingFunc(object)

	// ---- output binding value ---
	for _, bindingKey := range this.bindingKeys {
		this.setBinding(bindingKey, binding)
	}

}

func (this *baseBuilder) setBinding(bindingKey bindingKey, binding binding) {

	this.beanCtx.setBinding(bindingKey, binding)

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
