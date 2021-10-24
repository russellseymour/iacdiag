package models

type Config struct {
	Diagram Diagram `mapstructure:"diagram"`
	Log     Log     `mapstructure:"log"`
}
