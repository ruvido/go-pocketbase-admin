package letter

import (
	"os"
	"fmt"
	"bytes"
	"time"
	"strings"
	"io/ioutil"
	"github.com/yuin/goldmark"
	"github.com/adrg/frontmatter"
)


type Letter struct {
	Title   string    `yaml:"title"`
	Date    time.Time `yaml:"date"`
	Content string
}

func MarkdownToHtml( filename string ) Letter {

	fmt.Println("markdown")

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: loading markdown file "+filename)
	}

	var letter Letter
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &letter)
	if err != nil {
		fmt.Println("Error: reading the front matter")
	}

	// Abort if Subject is missing
	if letter.Title == "" {
		fmt.Fprintln(os.Stderr, "Subject is missing")
		os.Exit(99)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(rest, &buf); err != nil {
		panic(err)
	}

	letter.Content = buf.String()

	return letter
}
