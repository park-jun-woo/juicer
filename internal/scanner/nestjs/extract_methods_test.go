//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractMethods_Basic 테스트
package nestjs

import "testing"

func TestExtractMethods_Basic(t *testing.T) {
	src := []byte(`
@Controller('users')
export class UsersController {
  @Get()
  findAll() {}
  @Post()
  create() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no classes")
	}
	result := extractMethods(classes[0], src, "test.ts")
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
}
