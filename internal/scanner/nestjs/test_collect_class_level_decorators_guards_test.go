//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectClassLevelDecorators_Guards 테스트
package nestjs

import "testing"

func TestCollectClassLevelDecorators_Guards(t *testing.T) {
	src := []byte(`
@Controller('users')
@UseGuards(AuthGuard)
@Roles('admin')
export class UsersController {}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	cls := findAllByType(root, "class_declaration")[0]
	ci := &controllerInfo{}
	collectClassLevelDecorators(cls, src, ci)
	if len(ci.classMiddleware) == 0 {
		t.Fatalf("expected guards, got %v", ci.classMiddleware)
	}
	if len(ci.classRoles) == 0 {
		t.Fatalf("expected roles, got %v", ci.classRoles)
	}
}
