//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtendsFactoryName_NoExtends 테스트
package nestjs

import "testing"

func TestExtendsFactoryName_NoExtends(t *testing.T) {
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
	got := extendsFactoryName(classes[0], src)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
