//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractTypeParams_NoTypeParams 타입 파라미터 없는 클래스에서 빈 결과 반환 테스트
package nestjs

import "testing"

func TestExtractTypeParams_NoTypeParams(t *testing.T) {
	src := []byte(`
class PlainController {
  findAll() { }
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
	if len(got) != 0 {
		t.Fatalf("expected no type params, got %v", got)
	}
}
