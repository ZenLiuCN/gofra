package modeler

type Configurer int

func (c Configurer) IsModifier() bool {
	return int(c)&int(ConfigurerModifier) > 0
}
func (c Configurer) IsModified() bool {
	return int(c)&int(ConfigurerModified) > 0
}
func (c Configurer) IsSoftRemoved() bool {
	return int(c)&int(ConfigurerSoftRemoved) > 0
}
func (c Configurer) IsVersioned() bool {
	return int(c)&int(ConfigurerVersion) > 0
}
func MakeConfigurer(flags ...ConfigurerFlag) Configurer {
	c := 0
	for _, flag := range flags {
		c |= int(flag)
	}
	return Configurer(c)
}

type ConfigurerFlag int

var (
	ConfigurerAll = Configurer(0b00001111)
)

const (

	// ConfigurerModifier flag that requires record who took action
	ConfigurerModifier ConfigurerFlag = 1 << iota
	// ConfigurerModified flag that requires manually recording update timestamp
	ConfigurerModified
	// ConfigurerSoftRemoved flag that support soft remove via [Model].Delete
	ConfigurerSoftRemoved
	// ConfigurerVersion flag that support optimistic lock with Version field
	ConfigurerVersion
)
