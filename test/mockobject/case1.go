package mockobject

import "fmt"

type Case1Helper struct {
}

func (this *Case1Helper) SayHello() {
	fmt.Println("case 1 hello ")
}
