//ff:func feature=sql type=test control=sequence
//ff:what TestToQueryName_WithRepo 테스트
package sqls

import "testing"

func TestToQueryName_WithRepo(t *testing.T) {
	got := toQueryName("UserRepository.FindByID")
	if got != "UserFindByID" {
		t.Fatalf("expected UserFindByID, got %s", got)
	}
}

