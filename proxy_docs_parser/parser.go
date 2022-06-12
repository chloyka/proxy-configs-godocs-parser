package proxy_docs_parser

import (
	"go/doc"
	"go/parser"
	"go/token"
	"io/fs"
	"regexp"
	"strings"
)

var Host *string
var Port *int
var Scheme *string
var ProxyStructSlice []ProxyStruct

func Walk(path string, info fs.DirEntry, err error) error {
	if info.IsDir() && !strings.Contains(path, ".git") {
		if err != nil {
			panic(err)
		}
		fset := token.NewFileSet()
		d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		for _, f := range d {
			p := doc.New(f, "./", 0)
			for _, t := range p.Funcs {
				proxyObj := ProxyStruct{}
				backendObj := ProxyBackendStruct{}
				r := regexp.MustCompile(`(?P<Decorator>@[a-zA-Z]+)(?P<Delim>\s+)(?P<Value>["a-zA-Z/:\-_.]+)`)
				decoratorIndex := r.SubexpIndex("Decorator")
				valueIndex := r.SubexpIndex("Value")
				matches := r.FindAllStringSubmatch(t.Doc, -1)
				if len(matches) != 0 {
					for _, match := range matches {
						switch match[decoratorIndex] {
						case "@ExternalPath":
							proxyObj.Path = match[valueIndex]
						case "@ExternalMethod":
							proxyObj.Method = match[valueIndex]
						case "@InternalMethod":
							backendObj.Method = match[valueIndex]
						case "@InternalPath":
							backendObj.Path = match[valueIndex]
						}
					}
				}
				if backendObj != (ProxyBackendStruct{}) && proxyObj != (ProxyStruct{}) {
					backendObj.Host = *Host
					backendObj.Port = *Port
					backendObj.Scheme = *Scheme
					proxyObj.Backend = backendObj
					ProxyStructSlice = append(ProxyStructSlice, proxyObj)
				}
			}
		}
	}

	return nil
}
