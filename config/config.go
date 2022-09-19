package config

type Configuration struct {
	App       App       `mapstructure:"app" json:"app" yaml:"app"`
	Log       Log       `mapstructure:"log" json:"log" yaml:"log"`
	Databases Databases `mapstructure:"databases" json:"databases" yaml:"databases"`
	Jwt       Jwt       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Storage   Storage   `mapstructure:"storage" json:"storage" yaml:"storage"`
}
