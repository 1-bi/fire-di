package fire_di

import (
	"fmt"
	"reflect"
)

// ------- defined appp store ----
type BeanCtxBinder struct {
	bindingErrors []error
	bindings      map[bindingKey]binding
}

func (this *BeanCtxBinder) Bind(froms ...interface{}) Builder {
	// ---- create builder ----
	return this.bind(newBindingKey, froms)
}

// -------------------- bind context -------------
/**
 * bind function inteface with reflect type
 */
func (this *BeanCtxBinder) bind(newBindingKeyFunc func(reflect.Type) bindingKey, from []interface{}) InterfaceBuilder {

	lenFrom := len(from)
	if lenFrom == 0 {
		this.addBindingError(errNil)
		return newNoOpBuilder()
	}

	bindingKeys := make([]bindingKey, lenFrom)

	for i := 0; i < lenFrom; i++ {
		fromReflectType := reflect.TypeOf(from[i])

		if fromReflectType == nil {
			this.addBindingError(errNil)
			return newNoOpBuilder()
		}

		fromOrgType := fromReflectType.Elem()

		// --- customer array ---
		bindingKeys[i] = newBindingKeyFunc(fromOrgType)
	}

	return newBuilder(this, bindingKeys)
}

func (binderCtx *BeanCtxBinder) addBindingError(err error) {
	binderCtx.bindingErrors = append(binderCtx.bindingErrors, err)
}

/**
 * bind type in builder
 */
func (this *BeanCtxBinder) BindType(froms ...interface{}) Builder {

	// ---- create builder ----
	return this.bind(newBindingKey, froms)
}

/**
 * defined method binding
 * implement function for inject api
 */
func (this *BeanCtxBinder) String() string {
	return fmt.Sprintf("beanCtx{%s}", "update content ")
}

/**
 * create new beanCtx implement
 */
func createBeanCtxBinder() *BeanCtxBinder {
	return &BeanCtxBinder{make([]error, 0), make(map[bindingKey]binding)}
}

/**
 * set binding
 */
func (this *BeanCtxBinder) setBinding(bindingKey bindingKey, binding binding) {
	foundBinding, ok := this.bindings[bindingKey]
	if ok {
		this.addBindingError(errAlreadyBound.withTag("bindingKey", bindingKey).withTag("foundBinding", foundBinding))
		return
	}
	this.bindings[bindingKey] = binding
}
