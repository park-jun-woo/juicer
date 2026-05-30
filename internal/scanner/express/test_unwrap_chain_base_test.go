//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChain_Base 테스트
package express

import "testing"

func TestUnwrapChain_Base(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/:id').get(h);`))
	path, rv, methods := unwrapChain(outermostCall(fi), fi.Src, map[string]bool{"router": true})
	if path != "/:id" || rv != "router" || len(methods) != 1 {
		t.Fatalf("path=%q rv=%q methods=%v", path, rv, methods)
	}
}
