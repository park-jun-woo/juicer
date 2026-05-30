//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestMergeParentFields 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestMergeParentFields(t *testing.T) {
	got := mergeParentFields([]scanner.Field{{Name: "id"}, {Name: "name"}}, []scanner.Field{{Name: "name"}})
	if len(got) != 2 || got[0].Name != "id" {
		t.Fatalf("got %+v", got)
	}
}
