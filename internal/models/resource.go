package models

type Resource struct {
	Name      string     `mapstructure:"name"`
	Type      string     `mapstructure:"type"`
	Resources []Resource `mapstructure:"resources"`
}
