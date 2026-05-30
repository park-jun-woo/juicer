//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractOneRoute_NoFunctionDef 테스트
package flask

import "testing"

func TestExtractOneRoute_NoFunctionDef(t *testing.T) {

	src := `@register
class MyView:
    pass
`
	def, b := firstDecoratedDef(t, src)
	if routes := extractOneRoute(def, b, make(blueprintPrefix), "app.py"); routes != nil {
		t.Fatalf("decorated class should yield nil, got %v", routes)
	}
}
