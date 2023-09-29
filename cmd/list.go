package admin
import (

	"fmt"

	"github.com/ruvido/go-pocketbase-admin/pkg"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var searchFilter = ""

var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"l", "lis", "lst"},
	Short:  "List Records from a Pocketbase Collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := args[0]
		fmt.Println(collectionName)
		admin.ListRecordsFromCollection(collectionName, searchFilter)

		// template     := viper.GetString("general.template")
		// lett  := letter.MarkdownToHtml(markdownFile, template)
		// users := letter.FetchAddresses( list )
		// letter.Send ( lett, users ) // Send Newsletter
	},
}

func init() {
    listCmd.Flags().StringVarP(&searchFilter, "search", "s", "", "Search filter (e.g. name ~ 'chiara')")
    rootCmd.AddCommand(listCmd)
}
