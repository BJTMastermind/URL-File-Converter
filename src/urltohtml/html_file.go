package urltohtml

import (
    "fmt"
    "os"
    "strings"
)

type HTMLFile struct {
    Html string
}

func (html_file HTMLFile) WriteToHtmlFile(parsed_data Parser) {
    output_filename := strings.Split(parsed_data.Filename, ".")[0]+".html"

    output_data := []byte(parsed_data.Html_file.Html)
    err := os.WriteFile(output_filename, output_data, 0755)
    if err != nil {
        fmt.Printf("Failed to write html to %s.\n", output_filename)
    }
    fmt.Printf("Wrote html to %s successfully.\n", output_filename)
}