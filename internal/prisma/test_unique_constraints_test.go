//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what uniqueConstraints 필드 @unique 및 블록 @@unique 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestUniqueConstraints(t *testing.T) {
	s := schema{columns: map[string]map[string]string{"M": {"orgId": "org_id", "email": "email"}}}
	m := model{
		name:       "M",
		fields:     []field{{name: "email", attrs: []string{"@unique"}}},
		blockAttrs: []string{"@@index([orgId])", "@@unique([orgId, email])"},
	}
	got := uniqueConstraints(m, s)
	want := []string{`UNIQUE ("email")`, `UNIQUE ("org_id", "email")`}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := uniqueConstraints(model{name: "M"}, s); len(got) != 0 {
		t.Errorf("none: %v", got)
	}
}
