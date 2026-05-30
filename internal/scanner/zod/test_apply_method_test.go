//ff:func feature=scan type=test control=iteration dimension=1 topic=zod
//ff:what TestApplyMethod 테스트
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMethod(t *testing.T) {
	cases := []struct {
		name  string
		args  []string
		check func(f scanner.Field) bool
	}{
		{"string", nil, func(f scanner.Field) bool { return f.Type == "string" }},
		{"number", nil, func(f scanner.Field) bool { return f.Type == "number" }},
		{"int", nil, func(f scanner.Field) bool { return f.Type == "integer" }},
		{"boolean", nil, func(f scanner.Field) bool { return f.Type == "boolean" }},
		{"email", nil, func(f scanner.Field) bool { return f.Validate == "email" }},
		{"url", nil, func(f scanner.Field) bool { return f.Validate == "uri" }},
		{"uuid", nil, func(f scanner.Field) bool { return f.Validate == "uuid" }},
		{"optional", nil, func(f scanner.Field) bool { return f.Nullable }},
		{"nullable", nil, func(f scanner.Field) bool { return f.Nullable }},
		{"array", nil, func(f scanner.Field) bool { return f.Type == "array" }},
		{"object", nil, func(f scanner.Field) bool { return f.Type == "object" }},
		{"enum", []string{"a", "b"}, func(f scanner.Field) bool { return f.Type == "string" && len(f.Enum) == 2 }},
	}
	for _, c := range cases {
		f := scanner.Field{}
		ApplyMethod(&f, ChainMethod{Name: c.name, Args: c.args})
		if !c.check(f) {
			t.Errorf("ApplyMethod(%q) failed: %+v", c.name, f)
		}
	}
}
