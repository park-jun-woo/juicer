//ff:func feature=ddl type=test control=sequence
//ff:what buildFKDeps 테이블별 참조 집합 구성(외부참조/자기참조 제외) 테스트
package ddl

import "testing"

func TestBuildFKDeps(t *testing.T) {
	tables := map[string]*Table{
		"orgs":  {},
		"users": {Constraints: []string{`FOREIGN KEY (org_id) REFERENCES orgs (id)`}},
		"self":  {Constraints: []string{`FOREIGN KEY (p) REFERENCES self (id)`, `REFERENCES external (id)`}},
	}
	names := []string{"orgs", "users", "self"}
	deps := buildFKDeps(tables, names)
	if !deps["users"]["orgs"] {
		t.Errorf("users should depend on orgs: %v", deps["users"])
	}
	if len(deps["orgs"]) != 0 {
		t.Errorf("orgs no deps: %v", deps["orgs"])
	}
	// self-reference dropped, external dropped
	if len(deps["self"]) != 0 {
		t.Errorf("self should have no deps: %v", deps["self"])
	}
}
