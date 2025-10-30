/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	dataFile string
	cfgFile  string
	rootCmd  = &cobra.Command{
		Use:   "tri",
		Short: "A simple cli based todo app",
		Long: `tri helps you manage todos easily.

Configuration options:
  - Config file: $HOME/.tri.yaml (default)
  - Example contents:
      datafile: /home/user/code/tri/.tridos.json
  - Or set via environment variable:
      export TRI_DATAFILE=/home/user/code/tri/.tridos.json
  - Or use a flag:
      tri --datafile /path/to/file.json add "task"`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".tri" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tri")
	}

	viper.AutomaticEnv()      // read in environment variables that match
	viper.SetEnvPrefix("tri") // so TRI_DATAFILE works

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", viper.ConfigFileUsed())
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	home, err := homedir.Dir()

	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+"code/tri/.tridos.json", "file to store todos")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tri.yaml)")
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
