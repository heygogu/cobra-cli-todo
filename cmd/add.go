/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/heygogu/tri/todo"
)

// addCmd represents the add command
var (
	priority int
	addCmd   = &cobra.Command{
		Use:   "add",
		Short: "Add a todo",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: addRun,
	}
)

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	// fmt.Println(items)
	err = todo.SaveItems(viper.GetString("datafile"), items)

	if err != nil {
		fmt.Printf("%v", err)
	}
}
func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
