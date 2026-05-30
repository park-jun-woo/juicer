//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractOneRoute_Route 테스트
package flask

import "testing"

func TestExtractOneRoute_Route(t *testing.T) {
	src := `from flask import Flask
app = Flask(__name__)

@app.route('/users/<int:user_id>')
def get_user(user_id):
    return {}
`
	def, b := firstDecoratedDef(t, src)
	routes := extractOneRoute(def, b, make(blueprintPrefix), "app.py")
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].handler != "get_user" || routes[0].path != "/users/{user_id}" {
		t.Fatalf("route = %+v", routes[0])
	}
}
