package fire_di

import (
	"fmt"
)

type binding interface {
	fmt.Stringer
	// has to be a copy constructor
	resolvedBinding(*register, *injector) (resolvedBinding, error)
}

type resolvedBinding interface {
	fmt.Stringer
	validate() error
	get() (interface{}, error)
}

func newSingletonBinding(singleton interface{}) binding {
	return &singletonBinding{singleton, nil}
}

type singletonBinding struct {
	singleton interface{}
	injector  *injector
}

func (s *singletonBinding) String() string {
	return fmt.Sprintf("%v", s.singleton)
}

func (s *singletonBinding) validate() error {
	return nil
}

func (s *singletonBinding) get() (interface{}, error) {
	return s.singleton, nil
}

func (s *singletonBinding) resolvedBinding(provider *register, injector *injector) (resolvedBinding, error) {
	return &singletonBinding{s.singleton, injector}, nil
}
