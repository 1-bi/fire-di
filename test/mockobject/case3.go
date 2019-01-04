package mockobject

import "fmt"

type Case3MockObj1 struct {
}

func (this *Case3MockObj1) SayHello() {
	fmt.Println("case 3 mock object 1 ")
}

type Case3MockObj2 struct {
}

func (this *Case3MockObj2) SayHello() {
	fmt.Println("case 3 mock object 2 ")
}

type Case3MockObj3 struct {
}

func (this *Case3MockObj3) Inject(obj *Case3MockObj2) {
	fmt.Println(" run inject method ")

	// --- inject base boject
	obj.SayHello()
}

func (this *Case3MockObj3) SayHello() {
	fmt.Println("case 3 mock object 3 ")
}
