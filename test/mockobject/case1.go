package mockobject

import "fmt"

type Case1MockObj1 struct {
}

func (myself *Case1MockObj1) SayHello() {
	fmt.Println("case 1 mock object 1 ")
}

type Case1MockObj2 struct {
}

func (myself *Case1MockObj2) SayHello() {
	fmt.Println("case 1 mock object 2 ")
}
