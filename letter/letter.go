package letter

import (
	"github.com/ruvido/go-pocketbase-admin/pkg"
)

func SendEmails ( people []admin.User, filename string ) {

	letter := MarkdownToHtml( filename )
	Send ( letter, people ) 

}
