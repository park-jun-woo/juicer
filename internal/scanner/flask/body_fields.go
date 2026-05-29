//ff:type feature=scan type=model topic=flask
//ff:what Flask 핸들러 본문에서 추출한 바디 필드 모음
package flask

// bodyFields holds request body fields extracted from a Flask handler body.
type bodyFields struct {
	formFields []string // request.form[...] / request.form.get(...) keys
	jsonFields []string // request.json[...] / request.get_json()[...] keys
	hasJSON    bool     // request.json / request.get_json() referenced at all
}
