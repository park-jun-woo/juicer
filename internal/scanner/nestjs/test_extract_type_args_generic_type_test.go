//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractTypeArgs_GenericType 중첩 제네릭 타입 인자 추출 테스트
package nestjs

import (
	"strings"
	"testing"
)

func TestExtractTypeArgs_GenericType(t *testing.T) {
	src := []byte(`
@Controller('items')
export class ItemController extends BaseController<ResponseDto<Item>, ItemDto>() {
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
	if len(got) != 2 {
		t.Fatalf("expected 2 type args, got %d: %v", len(got), got)
	}
	if !strings.Contains(got[0], "ResponseDto") {
		t.Errorf("typeArgs[0] = %q, want something containing ResponseDto", got[0])
	}
	if got[1] != "ItemDto" {
		t.Errorf("typeArgs[1] = %q, want ItemDto", got[1])
	}
}
