//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestConvertFieldTypes 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestConvertFieldTypes(t *testing.T) {
	fields := []scanner.Field{{Name: "id", Type: "Long"}, {Name: "name", Type: "String"}}
	got := convertFieldTypes(fields)
	if got[0].Type != "integer" || got[1].Type != "string" {
		t.Fatalf("got %+v", got)
	}
}
