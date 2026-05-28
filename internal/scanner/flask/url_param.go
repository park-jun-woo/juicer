//ff:type feature=scan type=model topic=flask
//ff:what Flask URL 변수 파싱 결과 구조체
package flask

// urlParam holds a parsed Flask URL variable.
type urlParam struct {
	name      string
	converter string // "int", "float", "uuid", "path", "string", or ""
}
