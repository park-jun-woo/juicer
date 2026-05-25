//ff:type feature=scan type=model
//ff:what Endpoint 데이터 구조
package scanner

import (
	"go/ast"
)

// Endpoint — 발견된 HTTP 엔드포인트
type Endpoint struct {
	Method     string     `yaml:"method"               json:"method"`
	Path       string     `yaml:"path"                 json:"path"`
	Handler    string     `yaml:"handler"              json:"handler"`
	File       string     `yaml:"file,omitempty"       json:"file,omitempty"`
	Line       int        `yaml:"line,omitempty"       json:"line,omitempty"`
	Middleware []string   `yaml:"middleware,omitempty"  json:"middleware,omitempty"`
	Request    *Request   `yaml:"request,omitempty"     json:"request,omitempty"`
	Responses  []Response `yaml:"responses,omitempty"   json:"responses,omitempty"`

	// handlerExprs — handler.go가 함수 body를 해석할 때 사용 (직렬화 제외)
	handlerExprs []ast.Expr
}
