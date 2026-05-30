//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestSubstituteTypeParams 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestSubstituteTypeParams(t *testing.T) {
	fields := []scanner.Field{{Name: "data", Type: "T"}}
	got := substituteTypeParams(fields, []string{"T"}, []string{"string"})
	if got[0].Type != "string" {
		t.Fatalf("got %+v", got)
	}

	same := substituteTypeParams(fields, nil, nil)
	if same[0].Type != "T" {
		t.Fatalf("unchanged: %+v", same)
	}
}
