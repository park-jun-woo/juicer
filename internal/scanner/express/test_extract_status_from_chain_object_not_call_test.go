//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractStatusFromChain_ObjectNotCall 테스트
package express

import "testing"

func TestExtractStatusFromChain_ObjectNotCall(t *testing.T) {

	fi := mustParse(t, []byte(`res.json({});`))
	if got := extractStatusFromChain(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
