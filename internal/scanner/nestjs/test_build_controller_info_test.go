//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildControllerInfo 테스트
package nestjs

import "testing"

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
