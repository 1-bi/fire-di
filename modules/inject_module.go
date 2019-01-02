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
	this.delegate = this.createRouterDelegate(ctx)

}

/**
 * define base bean ---
 */
func (this *InjectSupportedModule) RegisterBean(bean interface{}) {

}

/**
 * =========================================================
 *  private function
 * =========================================================
 */
func (this *InjectSupportedModule) createRouterDelegate(ctx ModuleContext) *crouter.RouterDelegate {
	// ---- create router delegate -----
	delegate := crouter.NewRouterDelegate()

	// --- bind mapping delegate ----
	this.routerRegister(delegate)

	// --- inject bean for controller ---
	for _, handler := range delegate.AllRoulterPredefined() {

		bindingRef := handler.Ref

		refInjectObject, ok := bindingRef.(di.InjectAwaredSupport)

		if ok {

			injectFuns := refInjectObject.ProvideMethod()

			for _, infun := range injectFuns {

				ctx.GetProvider().InjectBean(infun)
			}

		}

	}

	return delegate
}
