package fire_di

import (
	"reflect"
)

type AppCtx struct {
	bindings map[bindingKey]resolvedBinding
}

/**
 * create new beanCtx implement
 */
func createAppCtx() *AppCtx {
	return &AppCtx{make(map[bindingKey]resolvedBinding)}
}

func (this *AppCtx) GetInstance(from interface{}) (interface{}, error) {
	return this.get(newBindingKey(reflect.TypeOf(from).Elem()))
}

func (this *AppCtx) get(bindingKey bindingKey) (interface{}, error) {
	binding, err := this.getBinding(bindingKey)
	if err != nil {
		return nil, err
	}
	return binding.get()
}

func (this *AppCtx) getBinding(bindingKey bindingKey) (resolvedBinding, error) {
	binding, ok := this.bindings[bindingKey]
	if !ok {
		return nil, errNoBinding.withTag("bindingKey", bindingKey)
	}
	return binding, nil
}
