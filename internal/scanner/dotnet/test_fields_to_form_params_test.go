//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestFieldsToFormParams 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestFieldsToFormParams(t *testing.T) {
	params := fieldsToFormParams([]scanner.Field{{Name: "a", Type: "string"}})
	if len(params) != 1 || params[0].Name != "a" {
		t.Fatalf("got %+v", params)
	}
}
