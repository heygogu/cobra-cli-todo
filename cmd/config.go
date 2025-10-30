package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Shows where your todos are stored and how configuration works",
	Long: `Display the current configuration details used by tri.

tri looks for your todo data file in this order:
  1. The '--datafile' flag (temporary override)
  2. Your ~/.tri.yaml config file (permanent preference)
  3. The 'TRI_DATAFILE' environment variable (session-level)
  4. The default path: $HOME/code/tri/.tridos.json

You can change where todos are saved using any of these methods.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuration in use:")
		fmt.Println("---------------------")

		// show config file location
		if file := viper.ConfigFileUsed(); file != "" {
			fmt.Println("Config file:", file)
		} else {
			fmt.Println("Config file: (none found)")
		}

		datafile := viper.GetString("datafile")
		fmt.Println("Data file:", datafile)
		fmt.Println()

		// tell user where it came from
		fmt.Print("Source: ")
		switch {
		case cmd.Flags().Changed("datafile"):
			fmt.Println("Flag (--datafile)")
		case viper.ConfigFileUsed() != "":
			fmt.Println("YAML config file (~/.tri.yaml)")
		case os.Getenv("TRI_DATAFILE") != "":
			fmt.Println("Environment variable (TRI_DATAFILE)")
		default:
			fmt.Println("Default built-in path")
		}

		fmt.Println("\nHelpful tips:")
		fmt.Println("  • To set it temporarily: tri list --datafile=/tmp/todos.json")
		fmt.Println("  • To set it permanently: echo 'datafile: /path/to/file.json' > ~/.tri.yaml")
		fmt.Println("  • To use an environment variable:")
		fmt.Println("      export TRI_DATAFILE=/path/to/file.json")
		fmt.Println("  • Default location (if nothing else set): $HOME/code/tri/.tridos.json")
		fmt.Println()
		fmt.Println("Run 'tri config' anytime to verify what’s currently in use.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
