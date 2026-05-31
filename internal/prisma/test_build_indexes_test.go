//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what buildIndexes @@index → CREATE INDEX 생성 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestBuildIndexes(t *testing.T) {
	s := schema{
		tableNames: map[string]string{"User": "users"},
		columns:    map[string]map[string]string{"User": {"orgId": "org_id"}},
	}
	m := model{name: "User", blockAttrs: []string{"@@index([orgId])", "@@unique([orgId])"}}
	got := buildIndexes(m, s)
	want := []string{`CREATE INDEX "users_org_id_idx" ON "users" ("org_id")`}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := buildIndexes(model{name: "User"}, s); len(got) != 0 {
		t.Errorf("no index: %v", got)
	}
}
