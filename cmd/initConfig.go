/*
Copyright © 2023 Nick Lesseos lesseosnick@gmail.com

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
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/blackflame007/codeview/colors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initConfigCmd represents the initConfig command
var initConfigCmd = &cobra.Command{
	Use:   "init-config",
	Short: "Initialize the configuration file with default colors",
	Long: `This command will create a default configuration file for codeview at the specified location,
populated with default color values.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define the path where the config file will be created
		configPath := filepath.Join(os.Getenv("HOME"), ".config", "codeview")
		configFile := filepath.Join(configPath, "codeview.yaml")

		// Check if the directory exists, if not, create it
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			os.MkdirAll(configPath, os.ModePerm)
		}

		// Open the file, creating it if it doesn't exist, with read-write permissions
		file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		defer file.Close()

		// Create a map to hold the colors
		colorsMap := make(map[string]string)

		// Populate the map with default colors
		for key, value := range colors.DefaultColors {
			// Enclose the keys in double quotes if they are not alphanumeric or underscore
			if regexp.MustCompile(`\W`).MatchString(key) {
				key = fmt.Sprintf("\"%s\"", key)
			}
			colorsMap[key] = value
		}

		// Marshal the colors map to YAML format
		colorsYAML, err := yaml.Marshal(map[string]interface{}{
			"colors": colorsMap,
		})
		if err != nil {
			fmt.Println("Error marshalling colors to YAML:", err)
			return
		}

		// Write the YAML data to the file
		_, err = file.Write(colorsYAML)
		if err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}

		fmt.Println("Config file initialized at:", configFile)
	},
}

func init() {
	rootCmd.AddCommand(initConfigCmd)
}
