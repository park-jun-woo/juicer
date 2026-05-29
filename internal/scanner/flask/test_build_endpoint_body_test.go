//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what 라우트 추출~buildEndpoint 변환에서 form/json requestBody 생성을 검증한다
package flask

import "testing"

func TestBuildEndpointBody(t *testing.T) {
	src := []byte(`from flask import Flask, request

app = Flask(__name__)

@app.route('/register', methods=['POST'])
def register():
    u = request.form['username']
    p = request.form.get('password')
    return ''

@app.route('/items', methods=['POST'])
def create_item():
    data = request.get_json()
    title = data['title']
    return ''

@app.route('/ping', methods=['POST'])
def ping():
    request.json
    return ''
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	routes := extractRoutes(root, src, make(blueprintPrefix), "app.py")

	got := map[string]routeInfo{}
	for _, r := range routes {
		got[r.handler] = r
	}

	reg := buildEndpoint(got["register"])
	if reg.Request == nil || len(reg.Request.FormFields) != 2 {
		t.Fatalf("register form fields = %+v", reg.Request)
	}

	item := buildEndpoint(got["create_item"])
	if item.Request == nil || item.Request.Body == nil || len(item.Request.Body.Fields) != 1 {
		t.Fatalf("create_item json body = %+v", item.Request)
	}
	if item.Request.Body.Method != "get_json" {
		t.Errorf("body method = %q, want get_json", item.Request.Body.Method)
	}

	ping := buildEndpoint(got["ping"])
	if ping.Request == nil || ping.Request.Body == nil {
		t.Fatalf("ping json body missing = %+v", ping.Request)
	}
	if len(ping.Request.Body.Fields) != 0 {
		t.Errorf("ping body fields = %v, want empty", ping.Request.Body.Fields)
	}
}
