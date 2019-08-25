package cmd

import "github.com/spf13/cobra"

var (
	cfgFile string
)

func initConfig() {
}

var rootCmd = &cobra.Command{
	Use:   "punch",
	Short: "punch",
}

// Execute runs the root command.
func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}
