package fire_di

import (
	"github.com/1-bi/log-api"
)

// BaseModuleContext define beanCtx object
type BaseModuleContext struct {
	Provider *register
}

// GetRegister get base register define
func (myself *BaseModuleContext) GetRegister() *register {
	return myself.Provider
}

/**
 * binding
 */
func registerModules(mods []Module) BaseModuleContext {

	// --- create new beanCtx  ----
	provider := newRegister()

	// --- bind common logger ----
	log := logapi.GetLogger("fire-di")

	if log != nil {

	}
	provider.loginst = log
	// ---- bind common logger ----

	// ---- create beanCtx context for beanCtx
	modCtx := BaseModuleContext{Provider: provider}

	/**
	 * define module interface
	 */
	for _, mod := range mods {
		// --- use and call function beanCtx ----
		mod.Bind(&modCtx)
	}

	// ---- return beanCtx context  ----
	return modCtx

}
