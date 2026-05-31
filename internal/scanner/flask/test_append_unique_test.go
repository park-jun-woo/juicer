//ff:func feature=scan type=test topic=flask control=sequence
//ff:what appendUnique 비빈/중복 제외 추가 테스트
package flask

import "testing"

func TestAppendUnique(t *testing.T) {
	list := appendUnique(nil, "a")
	list = appendUnique(list, "a") // dup
	list = appendUnique(list, "")  // empty
	list = appendUnique(list, "b")
	if len(list) != 2 || list[0] != "a" || list[1] != "b" {
		t.Errorf("got %v", list)
	}
}
