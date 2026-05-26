//ff:type feature=scan type=model topic=nestjs
//ff:what 메서드 파라미터 추출 결과 구조체
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// methodParams holds extracted parameter info from a method's formal parameters.
type methodParams struct {
	pathParams  []scanner.Param
	queryParams []scanner.Param
	bodyType    string
	files       []scanner.Param
}
