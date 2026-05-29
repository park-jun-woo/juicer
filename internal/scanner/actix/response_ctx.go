//ff:type feature=scan type=model topic=actix
//ff:what 응답 추출 시 공유되는 컨텍스트(블록·인덱스·캐시·누적 결과)
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

type responseCtx struct {
	block     *sitter.Node
	src       []byte
	sIdx      structIndex
	cache     map[string][]scanner.Field
	seen      map[string]bool
	responses []scanner.Response
}
