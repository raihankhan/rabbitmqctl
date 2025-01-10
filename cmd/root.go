package cmd

/*
Copyright © 2024 raihankhanraka@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rabbitmqctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("username", "u", "guest", "RabbitMQ User for authentication")
	rootCmd.PersistentFlags().StringP("password", "p", "guest", "RabbitMQ User's password for authentication")
	rootCmd.PersistentFlags().StringP("uri", "", "http://localhost:15672", "RabbitMQ cluster url")
	rootCmd.PersistentFlags().StringP("full-uri", "", "", "RabbitMQ cluster url with username and password")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rabbitmqctl",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: getHelp,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getHelp(cmd *cobra.Command, args []string) error {
	err := cmd.Help()
	if err != nil {
		fmt.Println("Failed to execute 'help' command. Reason:", err)
		return err
	}
	return nil
}
