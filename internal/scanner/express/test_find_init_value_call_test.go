//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindInitValue_Call 테스트
package express

import "testing"

func TestFindInitValue_Call(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	if v := findInitValue(firstDeclarator(t, fi)); v == nil || v.Type() != "call_expression" {
		t.Fatalf("got %v", v)
	}
}
