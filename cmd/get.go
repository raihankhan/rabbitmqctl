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
	"github.com/raihankhan/rabbitmqctl/pkg"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your application",
	Long:  `dgfg`,
	Run:   func(cmd *cobra.Command, args []string) {},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// rootCmd represents the base command when called without any subcommands
var getQueueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Get particular queue info or list of all queues",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		uri, _ := cmd.Flags().GetString("uri")
		client, err := pkg.GetClient(uri, username, password)
		if err != nil {
			return err
		}
		queues, err := pkg.GetQueues(client)
		if err != nil {
			return err
		}

		fmt.Println(queues)
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getQueueCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rabbitmqctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
