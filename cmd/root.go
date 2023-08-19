/*
Copyright © 2023 Alireza Arzehgar <alirezaarzehgar82@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "KafkaDataProcessingGo",
	Short: `Welcome to the Kafka Data Processing with Go project!
 This project showcases how to use Apache Kafka in combination with 
 the Go programming language to build a data processing application. `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
