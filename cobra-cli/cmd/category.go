/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"github.com/paulojr83/Go-Expert/cobra-cli/internal/database"
	"github.com/spf13/cobra"
	"log"
)

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("failed to open database: %v", err)
	}
	return db
}
func GetCategory(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*name, _ := cmd.Flags().GetString("name")
			exists, _ := cmd.Flags().GetBool("exists")
			id, _ := cmd.Flags().GetInt16("id")
			fmt.Println(category)
			fmt.Println(exists)
			fmt.Println(id)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Chamado antes do Run")
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Chamado depois do Run")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("Ocorreu um erro")
		},*/
	}}

//var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	/*
		categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Name of the category")
		//categoryCmd.PersistentFlags().StringP("name", "n", "", "Name of the category")
		categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists")
		categoryCmd.PersistentFlags().Int16P("id", "i", 0, "ID of the category")
	*/
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
