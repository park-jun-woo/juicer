//ff:type feature=scan type=model
//ff:what Endpoint 데이터 구조
package scanner

// Endpoint — 발견된 HTTP 엔드포인트
type Endpoint struct {
	Method     string     `yaml:"method"               json:"method"`
	Path       string     `yaml:"path"                 json:"path"`
	Handler    string     `yaml:"handler"              json:"handler"`
	File       string     `yaml:"file,omitempty"       json:"file,omitempty"`
	Line       int        `yaml:"line,omitempty"       json:"line,omitempty"`
	Middleware []string   `yaml:"middleware,omitempty"  json:"middleware,omitempty"`
	Roles      []string   `yaml:"roles,omitempty"      json:"roles,omitempty"`
	Request    *Request   `yaml:"request,omitempty"     json:"request,omitempty"`
	Responses  []Response `yaml:"responses,omitempty"   json:"responses,omitempty"`
}
