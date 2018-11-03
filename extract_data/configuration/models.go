package configuration

type MainConfig struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	ZwayServer string `yaml:"zway_Server,omitempty"`
	FollowedServices map[string]string `yaml:"followed_Services,omitempty"`
	ActivatedModules []string `yaml:"activated_Modules,omitempty"`
	DeviceTypes []string `yaml:"device_Types,omitempty"`
	DeviceConfiguration map[string]DeviceConf `yaml:"device_configuration,omitempty"`
}

type DeviceConf struct {
	Name string
	Room string
	Type string
	Unit string
	Ignore bool
}
