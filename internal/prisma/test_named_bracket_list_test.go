//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what namedBracketList 명명 인자의 괄호 토큰 추출 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestNamedBracketList(t *testing.T) {
	s := "fields: [authorId, orgId], references: [id]"
	if got := namedBracketList(s, "fields"); !reflect.DeepEqual(got, []string{"authorId", "orgId"}) {
		t.Errorf("fields: got %v", got)
	}
	if got := namedBracketList(s, "references"); !reflect.DeepEqual(got, []string{"id"}) {
		t.Errorf("references: got %v", got)
	}
	if got := namedBracketList(s, "missing"); got != nil {
		t.Errorf("missing name: %v", got)
	}
}
