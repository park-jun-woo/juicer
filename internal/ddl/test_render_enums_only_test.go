//ff:func feature=ddl type=test control=sequence
//ff:what TestRender_EnumsOnly 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestRender_EnumsOnly(t *testing.T) {
	enums := []EnumType{{Name: "status", Values: []string{"a"}}}
	out := Render(enums, map[string]*Table{})
	if !strings.Contains(out, "status") {
		t.Fatalf("expected enum status in output, got %q", out)
	}

	if strings.Contains(out, "CREATE TABLE") {
		t.Fatalf("expected no tables, got %q", out)
	}
}
