//ff:func feature=scan type=test control=sequence topic=express
//ff:what markChainProcessed: 체인 내 모든 call_expression 마킹
package express

import "testing"

func TestMarkChainProcessed(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x').get(h).put(h);`))
	processed := map[uintptr]bool{}
	markChainProcessed(outermostCall(fi), processed)

	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no calls")
	}
	for _, c := range calls {
		if !processed[uintptr(c.StartByte())] {
			t.Errorf("call at %d not marked processed", c.StartByte())
		}
	}
}
