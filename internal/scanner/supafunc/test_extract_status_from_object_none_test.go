//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractStatusFromObject_None 테스트
package supafunc

import "testing"

func TestExtractStatusFromObject_None(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { headers: {} };`))
	objs := findAllByType(fi.Root, "object")
	if got := extractStatusFromObject(objs[0], fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
