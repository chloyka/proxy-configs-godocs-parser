package main

import (
	"encoding/json"
	"flag"
	"fmt"
	docs_parser "main/docs-parser"
	"os"
	"path/filepath"
)

var OutputFile *string

func main() {
	docs_parser.Scheme = flag.String("scheme", "http", "service scheme")
	docs_parser.Host = flag.String("host", "", "service scheme")
	OutputFile = flag.String("outputFile", "proxy-config.json", "output file name")

	docs_parser.Port = flag.Int("port", 80, "service port")
	flag.Parse()

	if *docs_parser.Host == "" {
		fmt.Println("host cannot be empty")
		os.Exit(0)
	}

	_ = filepath.WalkDir(".", docs_parser.Walk)

	if len(docs_parser.ProxyStructSlice) > 0 {
		proxyJson, _ := json.Marshal(docs_parser.ProxyStructSlice)
		_ = os.WriteFile(*OutputFile, proxyJson, 655)
		os.Exit(0)
	}

	fmt.Println("Cannot find any proxy decorators. Please consult docs for troubleshooting")
	os.Exit(0)
}
