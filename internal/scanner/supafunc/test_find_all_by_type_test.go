//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestFindAllByType 테스트
package supafunc

import "testing"

func TestFindAllByType(t *testing.T) {
	fi := mustParse(t, []byte(`a(); b(); c();`))
	if len(findAllByType(fi.Root, "call_expression")) != 3 {
		t.Fatal("findAll")
	}
}
