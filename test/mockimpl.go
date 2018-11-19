package test

import "fmt"

type Testimpl struct {
}

type MockTestIIImpl struct {
	content string
}

/**
 * implement mock test
 */
func (impl *Testimpl) SayHello(content string) {

	fmt.Println(" OK hello .")
	fmt.Println(content)

}

func (impl *MockTestIIImpl) SayToMe(cont string) MockTestII {

	fmt.Println("call SayToMe ")

	testContent := MockTestIIImpl{content: cont}
	return MockTestII(&testContent)

}
