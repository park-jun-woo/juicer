//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what bracketList [...] 내부 콤마 토큰 추출 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestBracketList(t *testing.T) {
	if got := bracketList("[a, b , c]"); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("got %v", got)
	}
	if got := bracketList("noopen"); got != nil {
		t.Errorf("no open bracket: %v", got)
	}
	if got := bracketList("[unterminated"); got != nil {
		t.Errorf("no close bracket: %v", got)
	}
	if got := bracketList("[ , ]"); len(got) != 0 {
		t.Errorf("empty entries dropped: %v", got)
	}
}
