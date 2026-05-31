//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what resolveRawColumns 필드명 → 비인용 컬럼명(@map 반영/폴백) 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestResolveRawColumns(t *testing.T) {
	s := schema{columns: map[string]map[string]string{
		"User": {"userId": "user_id"},
	}}
	got := resolveRawColumns("User", []string{"userId", "unknown"}, s)
	if !reflect.DeepEqual(got, []string{"user_id", "unknown"}) {
		t.Errorf("got %v", got)
	}
	// model not in columns -> fallback to field names
	got = resolveRawColumns("Other", []string{"a"}, s)
	if !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("unknown model: %v", got)
	}
}
