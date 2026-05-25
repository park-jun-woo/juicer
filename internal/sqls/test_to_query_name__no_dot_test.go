//ff:func feature=sql type=parse control=sequence
//ff:what TestToQueryName_NoDot 테스트
package sqls

import "testing"

func TestToQueryName_NoDot(t *testing.T) {
	got := toQueryName("FindByID")
	if got != "FindByID" {
		t.Fatalf("expected FindByID, got %s", got)
	}
}
