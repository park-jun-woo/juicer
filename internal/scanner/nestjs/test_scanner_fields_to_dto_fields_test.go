//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestScannerFieldsToDTOFields 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScannerFieldsToDTOFields(t *testing.T) {
	fields := []scanner.Field{{Name: "a", Type: "string"}, {Name: "b", Type: "number"}}
	got := scannerFieldsToDTOFields(fields)
	if len(got) != 2 || got[0].name != "a" {
		t.Fatalf("got %+v", got)
	}
}
