//ff:func feature=scan type=test control=sequence
//ff:what TestConstToString_Int 테스트
package scanner

import (
	"go/constant"
	"testing"
)

func TestConstToString_Int(t *testing.T) {
	v := constant.MakeInt64(200)
	got := constToString(v)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}
