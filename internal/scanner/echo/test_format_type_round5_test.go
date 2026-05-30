//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestFormatType_Round5 테스트
package echo

import (
	"testing"
)

func TestFormatType_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
var A int
var B *int
var C []string
var D map[string]int
var E [3]byte
`)
	want := map[string]string{"A": "int", "B": "*int", "C": "[]string", "D": "map[string]int", "E": "[]byte"}
	for name, exp := range want {
		typ := defTypeByName(info, name)
		if typ == nil {
			t.Fatalf("no type for %s", name)
		}
		if got := formatType(typ); got != exp {
			t.Errorf("formatType(%s)=%q want %q", name, got, exp)
		}
	}
}
