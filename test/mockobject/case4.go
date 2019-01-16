package mockobject

import "fmt"

type Case4MockObj1 struct {
}

func (this *Case4MockObj1) SayHello() {
	fmt.Println("case 3 mock object 1 ")
}

/**
 * @Inject inject object
 */
func (this *Case4MockObj1) InjectMockObj(co *Case4MockObj2) {
	fmt.Println("  inject new object ")
	fmt.Println(co)
	co.SayHello()

}

type Case4MockObj2 struct {
}

func (this *Case4MockObj2) SayHello() {
	fmt.Println("case 3 mock object 2 ")
}

func (this *Case4MockObj2) Afterset() {
	fmt.Println("afterset case 3 mock object 2 ")
}
