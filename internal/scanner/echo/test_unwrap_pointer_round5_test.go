//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestUnwrapPointer_Round5 테스트
package echo

import (
	"go/types"
	"testing"
)

func TestUnwrapPointer_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
var P *int
var I int
`)
	var pt, it types.Type
	for id, obj := range info.Defs {
		if obj == nil {
			continue
		}
		if id.Name == "P" {
			pt = obj.Type()
		}
		if id.Name == "I" {
			it = obj.Type()
		}
	}
	if _, ok := unwrapPointer(pt).(*types.Basic); !ok {
		t.Fatalf("pointer not unwrapped: %v", unwrapPointer(pt))
	}
	if unwrapPointer(it) != it {
		t.Fatal("non-pointer should be returned unchanged")
	}
}
