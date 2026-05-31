//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveEnumPathArg_Fallback 테스트 (해석 불가 멤버표현식 원시 폴백)
package nestjs

import "testing"

// Unresolvable member expressions must fall back to the raw string (regression-safe).
func TestResolveEnumPathArg_Fallback(t *testing.T) {
	src := []byte(`
@Controller(Unknown.Thing)
export class C {
  @Get('x')
  m() {}
}
`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	ci, ok := buildControllerInfo(cls, src, "c.controller.ts", "/abs/c.controller.ts", map[string]string{}, root, "/tmp")
	if !ok {
		t.Fatal("expected ok")
	}
	if ci.prefix != "Unknown.Thing" {
		t.Fatalf("prefix fallback: want %q got %q", "Unknown.Thing", ci.prefix)
	}
}
