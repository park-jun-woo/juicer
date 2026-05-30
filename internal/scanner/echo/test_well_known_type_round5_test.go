//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestWellKnownType_Round5 테스트
package echo

import (
	"go/types"
	"testing"
)

func TestWellKnownType_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
import "time"
var T time.Time
`)
	typ := defTypeByName(info, "T")
	named, ok := typ.(*types.Named)
	if !ok {
		t.Fatal("no named type")
	}
	name, ok := wellKnownType(named)
	if !ok || name != "time.Time" {
		t.Fatalf("got %q %v", name, ok)
	}
}
