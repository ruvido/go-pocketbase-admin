package admin
import (

	"fmt"

	"github.com/ruvido/go-pocketbase-admin/pkg"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var backupSearchFilter = ""

var backupCmd = &cobra.Command{
	Use:   "backup",
	// Aliases: []string{"l", "lis", "lst"},
	Short:  "Backup Records from a Pocketbase Collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := args[0]
		fmt.Println(collectionName)
		admin.BackupCollection(collectionName, backupSearchFilter) 
	},
}

func init() {
    backupCmd.Flags().StringVarP(&backupSearchFilter, "search", "s", "", "Search filter (e.g. name ~ 'chiara')")
    rootCmd.AddCommand(backupCmd)
}
