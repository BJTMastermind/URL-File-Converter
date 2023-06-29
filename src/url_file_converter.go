package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/BJTMastermind/URL-File-Converter/urltohtml"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("error: Not enough arguments. Expected 1 got 0")
        return
    }
    if len(args) > 1 {
        fmt.Printf("error: Too many arguments. Expected 1 got %d\n", len(args))
        return
    }
    filepath := args[0]
    info, err := os.Stat(filepath)
    if err != nil {
        fmt.Println("error: Filename given doesn't exist or can't be read.")
        return
    }
    if info.IsDir() {
        fmt.Println("error: Filepath given is a directory.")
        return
    }
    if !strings.HasSuffix(filepath, ".url") {
        fmt.Println("error: Filename given is not a .url file.")
        return
    }

    parser := urltohtml.Parser{}
    parsed_data := parser.Parse(filepath)
    html := urltohtml.HTMLFile{}
    html.WriteToHtmlFile(parsed_data)
}
