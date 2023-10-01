package admin
import (

	"fmt"

	"github.com/ruvido/go-pocketbase-admin/pkg"
	"github.com/ruvido/go-pocketbase-admin/letter"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var sendSearchFilter = ""
var markdownFilename = ""


func init() {
    sendCmd.Flags().StringVarP(&markdownFilename, "markdown", "m", "", "Markdown file (email text)")
    sendCmd.Flags().StringVarP(&sendSearchFilter, "search", "s", "", "Search filter (e.g. name ~ 'chiara')")
    rootCmd.AddCommand(sendCmd)
}
var sendCmd = &cobra.Command{
	Use:   "send",
	// Aliases: []string{"l", "lis", "lst"},
	Short:  "Send a markdown email to a Pocketbase Collection (including filters)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := args[0]
		fmt.Println(collectionName)
		people := admin.CollectionRecords(collectionName, sendSearchFilter) 
		letter.SendEmails( people, markdownFilename)

		// template     := viper.GetString("general.template")
		// lett  := letter.MarkdownToHtml(markdownFile, template)
		// users := letter.FetchAddresses( list )
		// letter.Send ( lett, users ) // Send Newsletter
	},
}
