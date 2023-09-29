package admin
import (

	"fmt"

	"github.com/ruvido/go-pocketbase-admin/pkg"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var bubbleSearchFilter = ""

func init() {
    bubbleCmd.Flags().StringVarP(&bubbleSearchFilter, "search", "s", "", "Search filter (e.g. name ~ 'chiara')")
    rootCmd.AddCommand(bubbleCmd)
}

var bubbleCmd = &cobra.Command{
	Use:   "bubble",
	// Aliases: []string{"l", "lis", "lst"},
	Short:  "List Records from a Pocketbase Collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := args[0]
		fmt.Println(collectionName)
		// admin.CollectionRecords(collectionName, searchFilter) 
		people := admin.CollectionRecords(collectionName, bubbleSearchFilter) 
		admin.BubbleList(people)
	},
}


