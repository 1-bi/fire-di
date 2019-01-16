package mockobject

import "fmt"

type Case3MockObj1 struct {
}

func (myself *Case3MockObj1) SayHello() {
	fmt.Println("case 3 mock object 1 ")
}

/**
 * @Inject inject object
 */
func (myself *Case3MockObj1) InjectMockObj(co *Case3MockObj2, c3 *Case3MockObj3) {
	fmt.Println("  inject new object ")
	fmt.Println(co)
	co.SayHello()

	c3.SayHello()
}

type Case3MockObj2 struct {
}

func (myself *Case3MockObj2) SayHello() {
	fmt.Println("case 3 mock object 2 ")
}

func (myself *Case3MockObj2) Afterset() {
	fmt.Println("afterset case 3 mock object 2 ")
}

type Case3MockObj3 struct {
}

func (myself *Case3MockObj3) Inject(obj *Case3MockObj2) {
	fmt.Println(" run inject method ")

	// --- inject base boject
	obj.SayHello()
}

func (myself *Case3MockObj3) SayHello() {
	fmt.Println("case 3 mock object 3 ")
}

/**
 * define after method
 */
func (myself *Case3MockObj3) Afterset() {
	fmt.Println("call after set in case mock obj3")

}
