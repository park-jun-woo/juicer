//ff:func feature=sql type=parse control=sequence
//ff:what TestDetectDynamic 테스트
package sqls

import (
	"testing"
)

func TestDetectDynamic(t *testing.T) {
	t.Run("nil body", func(t *testing.T) {
		if detectDynamic(nil) {
			t.Error("expected false for nil body")
		}
	})

	t.Run("with += assignment", func(t *testing.T) {
		src := `package test
func f() {
	q := "SELECT * FROM users"
	q += " WHERE id = 1"
	_ = q
}
`
		body := parseMethodBody(t, src, "f")
		if !detectDynamic(body) {
			t.Error("expected true for += assignment")
		}
	})

	t.Run("with fmt.Sprintf", func(t *testing.T) {
		src := `package test
import "fmt"

func f() {
	q := fmt.Sprintf("SELECT * FROM %s", "users")
	_ = q
}
`
		body := parseMethodBody(t, src, "f")
		if !detectDynamic(body) {
			t.Error("expected true for fmt.Sprintf")
		}
	})

	t.Run("no dynamic patterns", func(t *testing.T) {
		src := `package test
func f() {
	q := "SELECT * FROM users"
	_ = q
}
`
		body := parseMethodBody(t, src, "f")
		if detectDynamic(body) {
			t.Error("expected false for static query")
		}
	})

	t.Run("non-fmt selector call", func(t *testing.T) {
		src := `package test
import "strings"

func f() {
	q := strings.ToUpper("select")
	_ = q
}
`
		body := parseMethodBody(t, src, "f")
		if detectDynamic(body) {
			t.Error("expected false for non-fmt call")
		}
	})
}
