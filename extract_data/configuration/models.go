package configuration

type MainConfig struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	FollowedServices map[string]string `yaml:"followedServices,omitempty"`
	ActivatedModules []string `yaml:"activatedModules,omitempty"`
}


