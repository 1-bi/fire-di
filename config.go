package fire_di

type baseConfigration struct {
	injectMethodPrefix []string
}

func (this *baseConfigration) getInjectMethodPrefix() []string {
	return this.injectMethodPrefix
}

/**
 * get the annotaion config
 */
type Configuration struct {
	baseConfigration
}

func (this *Configuration) SetInjectMethodPrefix(method ...string) {
	this.injectMethodPrefix = method
}
