//ff:type feature=scan type=model topic=fastapi
//ff:what 추출된 라우트 정보 구조체
package fastapi

import "github.com/park-jun-woo/juicer/internal/scanner"

// routeInfo holds information extracted from a decorated route function.
type routeInfo struct {
	method        string // HTTP method (GET, POST, etc.)
	path          string // route path from decorator
	handler       string // function name
	file          string // relative file path
	line          int    // line number (1-based)
	statusCode    int    // from status_code= in decorator
	responseModel string // from response_model= in decorator
	returnType    string // from -> annotation
	params        []scanner.Param
	query         []scanner.Param
	bodyType      string // Pydantic model type name
	files         []scanner.Param
	middleware    []string // Depends function names
}
