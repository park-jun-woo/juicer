//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectEnumElements 테스트
package nestjs

import "testing"

func TestCollectEnumElements(t *testing.T) {
	src := []byte(`const x = ['a', 'b', 1];`)
	root, _ := parseTypeScript(src)
	arr := findAllByType(root, "array")[0]
	vals := collectEnumElements(arr, src)
	if len(vals) != 3 || vals[0] != "a" || vals[2] != "1" {
		t.Fatalf("got %v", vals)
	}
}
