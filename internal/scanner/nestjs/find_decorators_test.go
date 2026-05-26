//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindDecorators_ExportedClass 테스트
package nestjs

import "testing"

func TestFindDecorators_ExportedClass(t *testing.T) {
	src := []byte(`
@Controller('users')
export class UsersController {}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no classes")
	}
	decs := findDecorators(classes[0], src)
	if len(decs) == 0 {
		t.Fatal("expected decorators")
	}
	if decs[0].name != "Controller" {
		t.Fatalf("expected Controller, got %q", decs[0].name)
	}
}
