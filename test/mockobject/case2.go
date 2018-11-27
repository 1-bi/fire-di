package mockobject

import "fmt"

type Case2MockObj1 struct {
}

func (this *Case2MockObj1) SayHello() {
	fmt.Println("case 2 mock object 1 ")
}

type Case2MockObj2 struct {
}

func (this *Case2MockObj2) SayHello() {
	fmt.Println("case 2 mock object 2 ")
}

type Case2MockObj3 struct {
}

func (this *Case2MockObj3) SayHello() {
	fmt.Println("case 2 mock object 3 ")
}
