package models

type Connection struct {
	Source string `mapstructure:"source"`
	Target string `mapstructure:"target"`
}
