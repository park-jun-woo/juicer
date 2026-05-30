//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestMergeParentFields 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestMergeParentFields(t *testing.T) {
	parent := []scanner.Field{{Name: "id"}, {Name: "name"}}
	own := []scanner.Field{{Name: "name"}, {Name: "email"}}
	got := mergeParentFields(parent, own)

	if len(got) != 3 {
		t.Fatalf("got %d: %+v", len(got), got)
	}
	if got[0].Name != "id" {
		t.Fatalf("first should be id: %+v", got)
	}
}
