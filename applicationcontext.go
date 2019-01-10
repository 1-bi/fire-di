package fire_di

import (
	"reflect"
)

// AppCtx define the application context
type AppCtx struct {
	bindings map[bindingKey]resolvedBinding
}

/**
 * create new beanCtx implement
 */
func createAppCtx() *AppCtx {
	return &AppCtx{make(map[bindingKey]resolvedBinding)}
}

// GetInstance the public instance for get application
func (myself *AppCtx) GetInstance(from interface{}) (interface{}, error) {
	return myself.get(newBindingKey(reflect.TypeOf(from).Elem()))
}

func (myself *AppCtx) get(bindingKey bindingKey) (interface{}, error) {
	binding, err := myself.getBinding(bindingKey)
	if err != nil {
		return nil, err
	}
	return binding.get()
}

func (myself *AppCtx) getBinding(bindingKey bindingKey) (resolvedBinding, error) {
	binding, ok := myself.bindings[bindingKey]
	if !ok {
		return nil, errNoBinding.withTag("bindingKey", bindingKey)
	}
	return binding, nil
}
