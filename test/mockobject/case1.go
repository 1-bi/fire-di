package mockobject

import "fmt"

type Case1MockObj1 struct {
}

func (this *Case1MockObj1) SayHello() {
	fmt.Println("case 1 mock object 1 ")
}

type Case1MockObj2 struct {
}

func (this *Case1MockObj2) SayHello() {
	fmt.Println("case 1 mock object 2 ")
}
