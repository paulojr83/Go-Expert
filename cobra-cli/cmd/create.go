/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/paulojr83/Go-Expert/cobra-cli/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command

func newCreateCmd(category database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Criar uma nova categoria",
		Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command. For example.`,
		RunE:  runCreateCategory(category),
	}
}
func runCreateCategory(db database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		_, err := db.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

var name string
var description string

func init() {
	createCmd := newCreateCmd(GetCategory(GetDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the category")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the category")
	/*createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("description")
	*/
	createCmd.MarkFlagsRequiredTogether("name", "description")

}
