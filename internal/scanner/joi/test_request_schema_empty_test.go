//ff:func feature=scan type=test topic=joi control=sequence
//ff:what RequestSchema.Empty 필드 유무 판정 테스트
package joi

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestRequestSchemaEmpty(t *testing.T) {
	if !(RequestSchema{}).Empty() {
		t.Error("zero value should be empty")
	}
	if (RequestSchema{Body: []scanner.Field{{Name: "x"}}}).Empty() {
		t.Error("with body should be non-empty")
	}
	if (RequestSchema{Query: []scanner.Field{{Name: "q"}}}).Empty() {
		t.Error("with query should be non-empty")
	}
}
