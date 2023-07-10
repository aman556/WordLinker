/*
Copyright © 2023 Aman Sharma
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "WordLinker",
	Short: "links the word to the respective link provided in the config file in the given file/dir.",
	Long: `WordLinker is a cli that links the words to the respective link provided in the config file in the text files/dir whose path is given at the time of runing the command.
	`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// func init() {
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.WordLinker.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
