package cmd

import (
	"log"

	"github.com/russellseymour/iacdiag/internal/models"
	"github.com/russellseymour/iacdiag/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Variable to hold the path to the configuration file
	cfgFile string

	// Set a varibale to set the version number of the app
	version string

	// App variable to hold information about how the application should work
	App models.App
)

var rootCmd = &cobra.Command{
	Use:     "iacdiag",
	Short:   "Generate infrastructure diagrams from code a configuration",
	Long:    "",
	Version: version,

	// Call prerun method to unmarshal the config into the app models
	PersistentPreRun: rootPreRun,
}

// Execute is the entry point for the application
func Execute() {

	// Determine if there is an error in the diagram
	err := rootCmd.Execute()

	if err != nil {
		log.Fatalf("%v", err)
	}
}

func init() {
	// Declare variables to accept the flag values
	var outputDir string

	outputDir = util.GetDefaultOutputDir()

	// Add flags that are used in every sub command
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to the configuration file to use")

	rootCmd.PersistentFlags().StringVarP(&outputDir, "outputdir", "o", outputDir, "Output directory for generated diagrams.")

	// Bind the command line arguments
	viper.BindPFlags(rootCmd.Flags())

}

// initConfig reads in a config file and ENV vars if set
// Also configures logging at the specified level
func initConfig() {

	if cfgFile != "" {
		// Use the config file from the cobra flag
		viper.SetConfigFile(cfgFile)
	}
}

func rootPreRun(ccmd *cobra.Command, args []string) {

	err := viper.Unmarshal(&App.Config)
	if err != nil {
		log.Fatalf("Unable to read configuration file: %v", err)
	}
}
