package admin
import (

	"fmt"

	"github.com/ruvido/go-pocketbase-admin/pkg"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var simpleSearchFilter = ""

func init() {
    simpleCmd.Flags().StringVarP(&simpleSearchFilter, "search", "s", "", "Search filter (e.g. name ~ 'chiara')")
    rootCmd.AddCommand(simpleCmd)
}

var simpleCmd = &cobra.Command{
	Use:   "simple",
	// Aliases: []string{"l", "lis", "lst"},
	Short:  "List Records from a Pocketbase Collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := args[0]
		fmt.Println(collectionName)
		// admin.CollectionRecords(collectionName, searchFilter) 
		people := admin.CollectionRecords(collectionName, simpleSearchFilter) 
		admin.SimpleList(people)
	},
}


