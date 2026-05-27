//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractTypeArgs_FiveParams 5개 제네릭 타입 인자 추출 테스트
package nestjs

import (
	"testing"
)

func TestExtractTypeArgs_FiveParams(t *testing.T) {
	src := []byte(`
@Controller('categories')
export class CategoryController extends BaseController<
  Category, CategoryDto, CreateCategoryDto, UpdateCategoryDto, FindAndSearchDto
>() {
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
	got := extractTypeArgs(classes[0], src)
	want := []string{"Category", "CategoryDto", "CreateCategoryDto", "UpdateCategoryDto", "FindAndSearchDto"}
	if len(got) != len(want) {
		t.Fatalf("expected %d type args, got %d: %v", len(want), len(got), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("typeArgs[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
