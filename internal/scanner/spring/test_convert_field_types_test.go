//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestConvertFieldTypes 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestConvertFieldTypes(t *testing.T) {
	got := convertFieldTypes([]scanner.Field{{Name: "id", Type: "Long"}})
	if got[0].Type != "integer" {
		t.Fatalf("got %+v", got)
	}
}
