//ff:func feature=ddl type=test control=sequence
//ff:what resolveRefKey REFERENCES 대상 → 테이블 키 매핑(정확/대소문자 폴백/없음) 테스트
package ddl

import "testing"

func TestResolveRefKey(t *testing.T) {
	tables := map[string]*Table{`"users"`: {}, "orgs": {}}
	// exact match
	if got := resolveRefKey(tables, `FOREIGN KEY (org_id) REFERENCES orgs (id)`); got != "orgs" {
		t.Errorf("exact: got %q", got)
	}
	// quote/case fallback: REFERENCES "Users" -> "users" key
	if got := resolveRefKey(tables, `FOREIGN KEY (x) REFERENCES "Users" (id)`); got != `"users"` {
		t.Errorf("fallback: got %q", got)
	}
	// no REFERENCES match
	if got := resolveRefKey(tables, `PRIMARY KEY (id)`); got != "" {
		t.Errorf("no ref: got %q", got)
	}
	// target absent from map
	if got := resolveRefKey(tables, `REFERENCES ghosts (id)`); got != "" {
		t.Errorf("absent: got %q", got)
	}
}
