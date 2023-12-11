package admin
import (

	"fmt"

	"github.com/ruvido/go-pocketbase-admin/pkg"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var dataFilename = ""

func init() {
    importCmd.Flags().StringVarP(&dataFilename, "data", "d", "", "Data file (e.g. csv)")
    rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import [collection]",
	// Aliases: []string{"l", "lis", "lst"},
	Short:  "Import Records from CSV to a Pocketbase Collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := args[0]
		fmt.Println(collectionName)
		fmt.Println(dataFilename)
		if len(dataFilename) == 0 {
			fmt.Println("Missing csv data file, e.g. '-d emails.csv'")
		} else {
			admin.ImportData(collectionName,dataFilename)
		}
		// people := admin.CollectionRecords(collectionName, bubbleSearchFilter) 
	},
}

// importCmd.SetHelpTemplate("Custom help template for the greet command")


