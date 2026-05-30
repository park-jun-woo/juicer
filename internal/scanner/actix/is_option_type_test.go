//ff:func feature=scan type=test control=sequence topic=actix
//ff:what isOptionType — Option<...> 판별을 검증
package actix

import "testing"

func TestIsOptionType(t *testing.T) {
	if !isOptionType("Option<String>") {
		t.Error("expected true for Option<String>")
	}
	if isOptionType("String") {
		t.Error("expected false for String")
	}
	if isOptionType("Vec<Option<i32>>") {
		t.Error("expected false: Option is not the outer type")
	}
}
