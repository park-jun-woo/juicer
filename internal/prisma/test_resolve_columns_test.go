//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what resolveColumns 필드명 → 인용 컬럼명 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestResolveColumns(t *testing.T) {
	s := schema{columns: map[string]map[string]string{
		"User": {"userId": "user_id"},
	}}
	got := resolveColumns("User", []string{"userId", "name"}, s)
	if !reflect.DeepEqual(got, []string{`"user_id"`, `"name"`}) {
		t.Errorf("got %v", got)
	}
}
