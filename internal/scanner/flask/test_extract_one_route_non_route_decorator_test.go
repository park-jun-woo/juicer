//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractOneRoute_NonRouteDecorator 테스트
package flask

import "testing"

func TestExtractOneRoute_NonRouteDecorator(t *testing.T) {
	src := `@staticmethod
def helper():
    pass
`
	def, b := firstDecoratedDef(t, src)
	if routes := extractOneRoute(def, b, make(blueprintPrefix), "app.py"); routes != nil {
		t.Fatalf("non-route decorator should yield nil, got %v", routes)
	}
}
