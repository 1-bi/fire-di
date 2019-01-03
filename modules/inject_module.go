package modules

import di "github.com/1-bi/fire-di"

type InjectSupportedModule struct {
	ctx di.ModuleContext
}

func (this *InjectSupportedModule) Bind(ctx di.ModuleContext) {

	/**
	 * create bind inject
	 */
	this.ctx = ctx

	// --- create delegate ----
	//this.delegate = this.createRouterDelegate(ctx)

}

/**
 * =========================================================
 *  private function
 * =========================================================
 */

/**
 * ---- define base bean ---
 */
func (this *InjectSupportedModule) RegisterBean(bean interface{}) {

}
