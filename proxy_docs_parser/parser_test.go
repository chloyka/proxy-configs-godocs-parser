package proxy_docs_parser

import (
	"path/filepath"
	"testing"
)

func TestParser(t *testing.T) {
	http := "http"
	host := "localhost"
	post := 80
	Scheme = &http
	Host = &host
	Port = &post

	_ = filepath.WalkDir(".", Walk)
	if len(ProxyStructSlice) < 1 {
		t.Error("Slice with proxy endpoints shouldn't be empty")
	}
	if ProxyStructSlice[0].Method != "GET" {
		t.Error("External http method should be defined as it describes in test mock")
	}
	if ProxyStructSlice[0].Backend.Method != "POST" {
		t.Error("Internal http method should be defined as it describes in test mock")
	}
	if ProxyStructSlice[0].Path != "/outside/path/to/endpoint.php" {
		t.Error("External http route should be defined as it describes in test mock")
	}
	if ProxyStructSlice[0].Backend.Path != "/internal/path/to/endpoint" {
		t.Error("Internal http route should be defined as it describes in test mock")
	}
}
