//ff:func feature=scan type=test topic=flask control=sequence
//ff:what collectCallFields/collectJSONVars/collectSubscriptFields 핸들러 본문 body 필드 수집 테스트
package flask

import "testing"

func TestCollectFields(t *testing.T) {
	code := `def handler():
    user = request.form.get('username')
    data = request.get_json()
    email = data['email']
    name = request.form['name']
`
	fn, src := flaskFuncDef(t, code)

	bf := collectCallFields(fn, src, bodyFields{})
	if !contains(bf.formFields, "username") {
		t.Errorf("form.get key not collected: %+v", bf)
	}
	if !bf.hasJSON {
		t.Error("get_json should mark hasJSON")
	}

	jsonVars := collectJSONVars(fn, src)
	if !jsonVars["data"] {
		t.Errorf("data should be json var: %v", jsonVars)
	}

	bf2 := collectSubscriptFields(fn, src, jsonVars, bodyFields{})
	if !contains(bf2.jsonFields, "email") {
		t.Errorf("json subscript key: %+v", bf2)
	}
	if !contains(bf2.formFields, "name") {
		t.Errorf("form subscript key: %+v", bf2)
	}
}
