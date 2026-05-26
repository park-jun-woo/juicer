//ff:type feature=scan type=model topic=nestjs
//ff:what 엔드포인트 추출 중간 결과 구조체
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// endpointInfo holds a single extracted endpoint before DTO resolution.
type endpointInfo struct {
	method     string
	path       string
	handler    string
	file       string
	line       int
	middleware []string
	roles      []string
	params     []scanner.Param
	query      []scanner.Param
	bodyType   string
	files      []scanner.Param
	statusCode int
	returnType string
}
