package mockobject

import "fmt"

type MockInterface interface {
	TestMock()
}

type SayHelloI interface {
	// SayHello define say hello
	SayHello(who string)
}

type SayHelloCase1 struct {
}

func (myself *SayHelloCase1) SayHello(who string) {

	fmt.Println("interface case 1 " + who)
}
