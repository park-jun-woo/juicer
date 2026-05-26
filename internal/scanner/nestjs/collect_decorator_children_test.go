//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what collectDecoratorChildren 테스트
package nestjs

import "testing"

func TestCollectDecoratorChildren_Found(t *testing.T) {
	src := []byte(`
@Controller('users')
export class UsersController {}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	exports := findAllByType(root, "export_statement")
	if len(exports) == 0 {
		t.Fatal("no exports")
	}
	decs := collectDecoratorChildren(exports[0], src)
	if len(decs) != 1 {
		t.Fatalf("expected 1, got %d", len(decs))
	}
}
