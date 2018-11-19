package test

/**
 * define logger factory handle
 */
type TestI interface {
	/**
	 * defined hello imple
	 */
	SayHello(content string)
}

type MockTestII interface {
	SayToMe(content string) MockTestII
}
