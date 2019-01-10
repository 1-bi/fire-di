package fire_di

/**
 * The api "inject" defined in this file.
 */

import (
	"fmt"
	"reflect"
)

/**
 * set config handle
 */
var runnimeConf = new(baseConfigration)

// Module define module interface
type Module interface {
	Bind(ctx ModuleContext)
}

// ModuleContext define module context
type ModuleContext interface {
	GetRegister() *register
}

// Binder create beanCtx interface
type Binder interface {
	fmt.Stringer

	// ---- bind provided handler for hag -----
	Provide(handlers ...interface{}) error

	// ---- set register bean ---
	InjectBean(funcs interface{}) error

	// ---- bind invoke handler
	Invoke(handlers ...interface{}) error
}

// CtxBinder create application inject
type CtxBinder interface {
	Bind(itypes ...interface{}) Builder
}

// ApplicationContext --- create application context ----
type ApplicationContext interface {

	// ---- get application method ----
	GetInstance(itypes ...interface{}) interface{}
}

// Builder is the return value from a Bind call from a Module.
type Builder interface {
	// --- bind singleton object ----
	ToProxyInst(singleton reflect.Value)
}

// InterfaceBuilder is the return value when binding an interface from a Module.
type InterfaceBuilder interface {
	Builder
}

// Injector is setting for injecto ioc handle
type Injector interface {

	/**
	 * create injector application
	 */
	Execute(funcs ...interface{}) error
}

// RegisterModules create function handle for beanCtx
func RegisterModules(mods ...Module) providerstore {

	// --- call beanCtx register function ----
	modContext := registerModules(mods)

	// ---- define beanCtx store ----
	bs := providerstore{modContext}

	return bs
}

// Config config the base inject environment
func Config(conf *Configuration) {

	inMethodPrefix := conf.baseConfigration.injectMethodPrefix

	lenMethodPrefix := len(inMethodPrefix)

	// --- use  default ---
	if lenMethodPrefix == 0 {
		runnimeConf.injectMethodPrefix = []string{"Inject"}
	}

}

// CreateInjector create injecto for object
func CreateInjector(bs providerstore) (*injector, error) {

	injector, err := createInjector(bs)

	return injector, err
}

// InjectAwaredSupport dev support api
type InjectAwaredSupport interface {

	// ProvideMethod defined provide method
	ProvideMethod() []interface{}
}
