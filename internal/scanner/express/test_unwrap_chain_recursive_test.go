//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChain_Recursive 테스트
package express

import "testing"

func TestUnwrapChain_Recursive(t *testing.T) {

	fi := mustParse(t, []byte(`router.route('/:id').get(getH).put(putH);`))
	path, rv, methods := unwrapChain(outermostCall(fi), fi.Src, map[string]bool{"router": true})
	if path != "/:id" || rv != "router" || len(methods) != 2 {
		t.Fatalf("path=%q rv=%q methods=%v", path, rv, methods)
	}
}
