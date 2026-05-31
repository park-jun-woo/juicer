//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what applyBodyFields 모든 routeInfo에 body 필드 주입 테스트
package flask

import "testing"

func TestApplyBodyFields(t *testing.T) {
	routes := []routeInfo{{method: "GET"}, {method: "POST"}}
	bf := bodyFields{formFields: []string{"u"}, jsonFields: []string{"e"}, hasJSON: true}
	out := applyBodyFields(routes, bf)
	for _, r := range out {
		if len(r.formFields) != 1 || len(r.jsonFields) != 1 || !r.hasJSONBody {
			t.Errorf("route not populated: %+v", r)
		}
	}
}
