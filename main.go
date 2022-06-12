package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chloyka/proxy-configs-godocs-parser/proxy_docs_parser"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var OutputFile *string

func main() {
	proxy_docs_parser.Scheme = flag.String("scheme", "http", "service scheme")
	proxy_docs_parser.Host = flag.String("host", "", "service scheme")
	OutputFile = flag.String("outputFile", "proxy-config.json", "output file name")

	proxy_docs_parser.Port = flag.Int("port", 80, "service port")
	flag.Parse()

	if *proxy_docs_parser.Host == "" {
		fmt.Println("host cannot be empty")
		os.Exit(0)
	}

	_ = filepath.WalkDir(".", proxy_docs_parser.Walk)

	if len(proxy_docs_parser.ProxyStructSlice) > 0 {
		proxyJson, _ := json.Marshal(proxy_docs_parser.ProxyStructSlice)
		err := ioutil.WriteFile(*OutputFile, proxyJson, 0644)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	fmt.Println("Cannot find any proxy decorators. Please consult docs for troubleshooting")
	os.Exit(0)
}
