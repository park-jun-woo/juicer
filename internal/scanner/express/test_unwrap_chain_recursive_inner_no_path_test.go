//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChainRecursive_InnerNoPath 테스트
package express

import "testing"

func TestUnwrapChainRecursive_InnerNoPath(t *testing.T) {

	fi := mustParse(t, []byte(`foo().put(p);`))
	inner, outer := recursiveParts(t, fi)
	if path, _, m := unwrapChainRecursive(inner, outer, "put", fi.Src, map[string]bool{"router": true}); path != "" || m != nil {
		t.Fatalf("expected empty, got %q %v", path, m)
	}
}
