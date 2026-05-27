//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractTypeParams 제네릭 타입 파라미터 추출 테스트
package nestjs

import "testing"

func TestExtractTypeParams(t *testing.T) {
	src := []byte(`
class GenericsController<D, B, CreateDtoType, UpdateDtoType, FindAndSearchDto> {
  findAll(): D[] { }
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
	got := extractTypeParams(classes[0], src)
	want := []string{"D", "B", "CreateDtoType", "UpdateDtoType", "FindAndSearchDto"}
	if len(got) != len(want) {
		t.Fatalf("expected %d type params, got %d: %v", len(want), len(got), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("typeParams[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
