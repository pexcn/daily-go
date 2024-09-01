package cmd

import (
	"daily/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "daily",
	Short: "A daily list generator",
	Long:  `A daily list generator, written in Go. 🚀`,
	// TODO: get version from build date
	Version: "0.1",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var globalFlag = &config.GlobalFlag{}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&globalFlag.Verbose, "verbose", "v", false, "Verbose output")

	//rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().SortFlags = false
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
