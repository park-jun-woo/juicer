//ff:type feature=scan type=model topic=flask
//ff:what 추출된 Flask 라우트 정보 구조체
package flask

// routeInfo holds information extracted from a single Flask route.
type routeInfo struct {
	method  string // HTTP method (GET, POST, etc.)
	path    string // resolved path (with blueprint prefix, OpenAPI format)
	handler string // function name
	file    string // relative file path
	line    int    // line number (1-based)
	params  []urlParam

	formFields  []string // request.form access keys (multipart/form-data)
	jsonFields  []string // request.json / get_json access keys (application/json)
	hasJSONBody bool     // request.json / request.get_json() referenced at all
}
