//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestFieldsToFormParams 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestFieldsToFormParams(t *testing.T) {
	fields := []scanner.Field{
		{Name: "name", Type: "string"},
		{Name: "userName", JSON: "user_name", Type: "string"},
	}
	params := fieldsToFormParams(fields)
	if len(params) != 2 || params[1].Name != "user_name" {
		t.Fatalf("got %+v", params)
	}
}
