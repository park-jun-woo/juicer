//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractorTestSetup 테스트 헬퍼
package actix

import "github.com/park-jun-woo/codistill/internal/scanner"

func extractorTestSetup() (*scanner.Endpoint, structIndex, map[string][]scanner.Field) {
	return &scanner.Endpoint{}, structIndex{}, map[string][]scanner.Field{}
}
