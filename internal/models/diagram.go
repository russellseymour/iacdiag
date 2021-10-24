package models

type Diagram struct {
	ProjectName string `mapstructure:"project_name"` // Name of the project
	DiagramName string `mapstructure:"diagram_name"` // Name of the diagram

	Resources   []Resource   `mapstructure:"resources"`   // List of cloud resources
	Connections []Connection `mapstructure:"connections"` // How all the components are connected

}
