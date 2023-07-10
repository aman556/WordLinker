/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/aman556/WordLinker/util"
	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "link the words present in the file/dir",
	Long:  `link command will converte the words into the links by refering to the config file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var DirPath string

		DirPath = args[0]
		err := util.LinkWords(DirPath)

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(linkCmd)
}
