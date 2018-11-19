package modules

import (
	"fmt"
	di "github.com/1-bi/fire-di"
)

type TestCaeeModule struct {
}

func (this *TestCaeeModule) Bind(ctx di.ModuleContext) {

	fmt.Println("Debug module ")

}
