package mockobject

import "fmt"

type Case2MockObj1 struct {
}

func (myself *Case2MockObj1) SayHello() {
	fmt.Println("case 2 mock object 1 ")
}

type Case2MockObj2 struct {
}

func (myself *Case2MockObj2) SayHello() {
	fmt.Println("case 2 mock object 2 ")
}

func (myself *Case2MockObj2) TestMock() {

	fmt.Println(" test mock interface ")

}

type Case2MockObj3 struct {
}

func (myself *Case2MockObj3) SayHello() {
	fmt.Println("case 2 mock object 3 ")
}
