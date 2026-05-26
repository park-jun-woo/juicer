//ff:func feature=sql type=test control=sequence
//ff:what TestToQueryName_NoDotCov 테스트
package sqls

import "testing"

func TestToQueryName_NoDotCov(t *testing.T) {
	got := toQueryName("NoDot")
	if got != "NoDot" {
		t.Fatalf("expected NoDot, got %s", got)
	}
}
