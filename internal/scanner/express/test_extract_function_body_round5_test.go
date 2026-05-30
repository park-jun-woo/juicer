//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractFunctionBody_Round5 테스트
package express

import "testing"

func TestExtractFunctionBody_Round5(t *testing.T) {

	fi := mustParse(t, []byte(`const f = () => { return 1; };`))
	arrow := exFirst(t, fi, "arrow_function")
	body := extractFunctionBody(arrow)
	if body == nil || body.Type() != "statement_block" {
		t.Fatalf("arrow body: %v", body)
	}

	fi2 := mustParse(t, []byte(`const g = function () { return 2; };`))
	fexpr := exFirst(t, fi2, "function_expression")
	body2 := extractFunctionBody(fexpr)
	if body2 == nil || body2.Type() != "statement_block" {
		t.Fatalf("function_expression body: %v", body2)
	}
}
