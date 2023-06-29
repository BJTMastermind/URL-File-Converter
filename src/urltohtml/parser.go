package urltohtml

import (
    "fmt"
    "log"
    "os"
    "strings"
)

type Parser struct {
    Filename string
    Url_file URLFile
    Html_file HTMLFile
}

func (parser Parser) Parse(filename string) Parser {
    content, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    lines := strings.Split(string(content), "\n")
    url := strings.Split(lines[1], "=")[1]

    parser.Filename = filename
    parser.Url_file.Url = url
    parser.Html_file.Html = fmt.Sprintf("<!DOCTYPE html>\n"+
    "<html>\n"+
    "    <head>\n"+
    "        <title>Internet Shortcut</title>\n"+
    "        <meta http-equiv=\"refresh\" content=\"0; url=%s\"/>\n"+
    "    </head>\n"+
    "    <body>\n"+
    "        <p>Redirecting to <a href=\"%s\">%s</a></p>\n"+
    "    </body>\n"+
    "</html>", url, url, url)
    return parser
}