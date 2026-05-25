//ff:func feature=scan type=extract control=sequence
//ff:what TestConstToString 테스트
package scanner

import (
	"go/constant"
	"testing"
)

func TestConstToString(t *testing.T) {
	t.Run("int constant", func(t *testing.T) {
		v := constant.MakeInt64(200)
		got := constToString(v)
		if got != "200" {
			t.Errorf("expected '200', got %q", got)
		}
	})

	t.Run("string constant", func(t *testing.T) {
		v := constant.MakeString("hello")
		got := constToString(v)
		if got != `"hello"` {
			t.Errorf("expected '\"hello\"', got %q", got)
		}
	})

	t.Run("float constant", func(t *testing.T) {
		v := constant.MakeFloat64(3.14)
		got := constToString(v)
		// Should use ExactString
		if got == "" {
			t.Error("expected non-empty string")
		}
	})

	t.Run("bool constant", func(t *testing.T) {
		v := constant.MakeBool(true)
		got := constToString(v)
		if got != "true" {
			t.Errorf("expected 'true', got %q", got)
		}
	})
}
