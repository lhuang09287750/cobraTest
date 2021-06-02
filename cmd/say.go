/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*		if len(args) < 1 {
					fmt.Println("Hello World!")
				}else {
					for _,item :=range args {
						fmt.Println("Hello", item, "!")
					}
				}*/

		/*		name, _:= cmd.Flags().GetString("name")
				if name == "" {
					name = "World"
				}
				fmt.Println("Hello " + name)*/

		greeting := "Hello"
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}

		if viper.GetString("name") != "" {
			name = viper.GetString("name")
		}
		if viper.GetString("greeting") != "" {
			greeting = viper.GetString("greeting")
		}
		fmt.Println(greeting + " " + name)

	},
}

func init() {
	rootCmd.AddCommand(sayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	sayCmd.Flags().StringP("name", "n", viper.GetString("ENVNAME"), "Set your name")
}
