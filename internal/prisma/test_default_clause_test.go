//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what defaultClause @default 존재여부 및 SQL 표현 변환 테스트
package prisma

import "testing"

func TestDefaultClause(t *testing.T) {
	// no @default
	if v, ok := defaultClause(field{}); ok || v != "" {
		t.Errorf("no default: got (%q,%v)", v, ok)
	}
	// autoincrement -> ("", true)
	if v, ok := defaultClause(field{attrs: []string{"@default(autoincrement())"}}); !ok || v != "" {
		t.Errorf("autoincrement: got (%q,%v), want (\"\",true)", v, ok)
	}
	// now() -> ("now()", true)
	if v, ok := defaultClause(field{attrs: []string{"@id", "@default(now())"}}); !ok || v != "now()" {
		t.Errorf("now(): got (%q,%v)", v, ok)
	}
}
