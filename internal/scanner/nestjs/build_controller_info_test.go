//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what buildControllerInfo / parseTypeScript / resolveDTOExtends / resolveParentDTOFields 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestBuildControllerInfo(t *testing.T) {
	src := []byte(`
@Controller('users')
export class UsersController {
  @Get(':id')
  findOne() {}
}
`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	ci, ok := buildControllerInfo(cls, src, "users.controller.ts", "/abs/users.controller.ts", map[string]string{})
	if !ok {
		t.Fatal("expected ok")
	}
	if ci.prefix != "users" {
		t.Fatalf("prefix: %q", ci.prefix)
	}
	if len(ci.endpoints) == 0 {
		t.Fatalf("expected endpoints, got %+v", ci.endpoints)
	}
}

func TestBuildControllerInfo_NotController(t *testing.T) {
	src := []byte(`export class PlainClass {}`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	if _, ok := buildControllerInfo(cls, src, "x.ts", "/abs/x.ts", map[string]string{}); ok {
		t.Fatal("expected not ok for non-controller")
	}
}

func TestParseTypeScript(t *testing.T) {
	root, err := parseTypeScript([]byte(`const x = 1;`))
	if err != nil || root == nil {
		t.Fatalf("err: %v", err)
	}
}

func TestResolveDTOExtends_NoHeritage(t *testing.T) {
	src := []byte(`class Dto { name: string; }`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	if got := resolveDTOExtends(cls, src, "x.ts", map[string]string{}, "", map[string][]scanner.Field{}); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestResolveDTOExtends_PlainExtends(t *testing.T) {
	src := []byte(`class Child extends Parent {}`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	// Parent not in imports -> resolveParentDTO returns empty -> nil dtoFields
	got := resolveDTOExtends(cls, src, "x.ts", map[string]string{}, "", map[string][]scanner.Field{})
	if len(got) != 0 {
		t.Fatalf("expected empty, got %+v", got)
	}
}

func TestResolveParentDTOFields_Cached(t *testing.T) {
	cache := map[string][]scanner.Field{
		"ParentDto": {{Name: "x", Type: "string"}},
	}
	got := resolveParentDTOFields("ParentDto", "x.ts", map[string]string{}, "", cache)
	if len(got) != 1 || got[0].Name != "x" {
		t.Fatalf("got %+v", got)
	}
}

func TestResolveParentDTOFields_NoImport(t *testing.T) {
	got := resolveParentDTOFields("Unknown", "x.ts", map[string]string{}, "", map[string][]scanner.Field{})
	if got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}
