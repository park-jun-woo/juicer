//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressCall_NotCall 테스트
package express

import "testing"

func TestIsExpressCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const a = express;`))

	ids := findAllByType(fi.Root, "identifier")
	if len(ids) == 0 {
		t.Fatal("no identifier")
	}
	if isExpressCall(ids[0], fi.Src) {
		t.Fatal("expected false for non-call")
	}
}
