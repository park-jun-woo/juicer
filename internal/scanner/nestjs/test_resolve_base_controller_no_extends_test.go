//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveBaseController_NoExtends 테스트
package nestjs

import "testing"

func TestResolveBaseController_NoExtends(t *testing.T) {
	src := []byte(`
@Controller('users')
export class UsersController {
  @Get()
  findAll() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no classes found")
	}
	imports := extractImports(root, src)
	eps := resolveBaseController(classes[0], src, "/tmp/test.ts", imports, "test.ts")
	if len(eps) != 0 {
		t.Fatalf("expected 0, got %d", len(eps))
	}
}
