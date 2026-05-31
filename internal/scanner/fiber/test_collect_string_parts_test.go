//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what collectStringParts 리터럴/문자열연결에서 문자열 조각 수집(변수/필드 스킵) 테스트
package fiber

import (
	"go/parser"
	"reflect"
	"testing"
)

func TestCollectStringParts(t *testing.T) {
	expr, err := parser.ParseExpr(`"a" + b + "c" + opts.Base`)
	if err != nil {
		t.Fatal(err)
	}
	var parts []string
	collectStringParts(expr, &parts)
	if !reflect.DeepEqual(parts, []string{"a", "c"}) {
		t.Errorf("got %v (variable/selector parts must be skipped)", parts)
	}
}
