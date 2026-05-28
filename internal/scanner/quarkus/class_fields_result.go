//ff:type feature=scan type=model topic=quarkus
//ff:what 클래스 필드 해석 결과 구조체
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

type classFieldsResult struct {
	fields     []scanner.Field
	typeParams []string
}
