//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractTypeArgs_NoExtends extends 없는 클래스에서 빈 결과 반환 테스트
package nestjs

import "testing"

func TestExtractTypeArgs_NoExtends(t *testing.T) {
	src := []byte(`
@Controller('foo')
export class FooController {
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
	if len(got) != 0 {
		t.Fatalf("expected no type args, got %v", got)
	}
}
