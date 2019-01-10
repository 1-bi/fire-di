package fire_di

type baseConfigration struct {
	injectMethodPrefix []string
}

func (myself *baseConfigration) getInjectMethodPrefix() []string {
	return myself.injectMethodPrefix
}

// Configuration get the annotaion config
type Configuration struct {
	baseConfigration
}

// SetInjectMethodPrefix set inject method prefix
func (myself *Configuration) SetInjectMethodPrefix(method ...string) {
	myself.injectMethodPrefix = method
}
