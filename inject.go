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
var annotation = false

/**
 * define module interface
 */
type Module interface {
	Bind(ctx ModuleContext)
}

/**
 * define module context
 */
type ModuleContext interface {
	GetRegister() *register
}

/**
 * create beanCtx interface
 */
type Binder interface {
	fmt.Stringer

	// ---- bind provided handler for hag -----
	Provide(handlers ...interface{}) error

	// ---- set register bean ---
	InjectBean(funcs interface{}) error

	// ---- bind invoke handler
	Invoke(handlers ...interface{}) error
}

/**
 * create application inject
 */
type CtxBinder interface {
	Bind(itypes ...interface{}) Builder
}

/**
 * --- create application context ----
 */
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

/**
 * create function handle for beanCtx
 */
func RegisterModules(mods ...Module) providerstore {

	// --- call beanCtx register function ----
	modContext := registerModules(mods)

	// ---- define beanCtx store ----
	bs := providerstore{modContext}

	return bs
}

/**
 * config the base inject environment
 */
func Config(annoConf *AnnotationConfig) {

	annotation = annoConf.AnnotationSupport

}

/**
 * create injecto for object
 */
func CreateInjector(bs providerstore) (*injector, error) {

	injector, err := createInjector(bs)

	return injector, err
}

/**
 * dev support api
 */
type InjectAwaredSupport interface {

	/**
	 * defined provide method
	 */
	ProvideMethod() []interface{}
}
