package modules

import (
	"fmt"
	"github.com/1-bi/fire-di/di"
)

type TestCaeeModule struct {
}

func (this *TestCaeeModule) Bind(ctx di.ModuleContext) {

	fmt.Println("Debug module ")

}
