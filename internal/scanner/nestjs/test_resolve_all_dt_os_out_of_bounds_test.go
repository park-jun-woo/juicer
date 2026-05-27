//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveAllDTOs_OutOfBounds 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveAllDTOs_OutOfBounds(t *testing.T) {
	reqs := []dtoRequest{{typeName: "Dto", epIdx: 100}}
	eps := []scanner.Endpoint{{Method: "GET"}}
	resolveAllDTOs(reqs, eps)
}
