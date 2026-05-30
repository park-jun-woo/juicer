//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestMergeParentFields_NoParent 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestMergeParentFields_NoParent(t *testing.T) {
	own := []scanner.Field{{Name: "x"}}
	got := mergeParentFields(nil, own)
	if len(got) != 1 {
		t.Fatalf("got %+v", got)
	}
}
