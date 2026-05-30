//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChain_NotCall 테스트
package express

import "testing"

func TestUnwrapChain_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`router;`))
	ids := findAllByType(fi.Root, "identifier")
	if p, _, m := unwrapChain(ids[0], fi.Src, map[string]bool{"router": true}); p != "" || m != nil {
		t.Fatalf("expected empty")
	}
}
