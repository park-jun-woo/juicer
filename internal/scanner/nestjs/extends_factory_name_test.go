//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtendsFactoryName_FactoryCall 테스트
package nestjs

import "testing"

func TestExtendsFactoryName_FactoryCall(t *testing.T) {
	src := []byte(`
@Controller('categories')
export class CategoryController extends BaseController<Category, CategoryDto>(CreateDto, UpdateDto) {
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
	if got != "BaseController" {
		t.Fatalf("expected BaseController, got %q", got)
	}
}
