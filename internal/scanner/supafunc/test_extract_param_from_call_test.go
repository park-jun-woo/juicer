//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what TestExtractParamFromCall 테스트
package supafunc

import "testing"

func TestExtractParamFromCall(t *testing.T) {
	fi := mustParse(t, []byte(`url.searchParams.get("limit");`))
	calls := findAllByType(fi.Root, "call_expression")
	for _, c := range calls {
		if got := extractParamFromCall(c, fi.Src); got == "limit" {
			return
		}
	}
	t.Fatal("did not find limit")
}
