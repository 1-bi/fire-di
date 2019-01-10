package fire_di

import (
	"gitlab.com/1-bi/log-api/loggerzap"
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
	provider := createProvider()

	// --- bind common logger ----
	log := loggerzap.GetLogger()

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
