package types

type ProxyStruct struct {
	Path    string             `json:"path"`
	Method  string             `json:"method"`
	Backend ProxyBackendStruct `json:"backend"`
}

type ProxyBackendStruct struct {
	Host   string `json:"host"`
	Scheme string `json:"scheme"`
	Port   int    `json:"port"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
