//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what collectArrayStringArgs 배열 노드의 문자열 요소 수집 테스트
package nestjs

import (
	"reflect"
	"testing"
)

func TestCollectArrayStringArgs(t *testing.T) {
	src := []byte("const a = ['/x', '/y', 3];")
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	arr := findAllByType(root, "array")[0]
	got := collectArrayStringArgs(arr, src)
	if !reflect.DeepEqual(got, []string{"/x", "/y"}) {
		t.Errorf("got %v (numbers must be skipped)", got)
	}
}
