//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestFieldsToFormParams 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestFieldsToFormParams(t *testing.T) {
	params := fieldsToFormParams([]scanner.Field{{Name: "a", Type: "string"}, {Name: "b", JSON: "bb", Type: "string"}})
	if len(params) != 2 || params[1].Name != "bb" {
		t.Fatalf("got %+v", params)
	}
}
