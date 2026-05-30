//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractStatusFromObject 테스트
package supafunc

import "testing"

func TestExtractStatusFromObject(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { status: 201, headers: {} };`))
	objs := findAllByType(fi.Root, "object")
	if got := extractStatusFromObject(objs[0], fi.Src); got != "201" {
		t.Fatalf("got %q", got)
	}
}
