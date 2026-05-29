//ff:func feature=scan type=test control=sequence topic=flask
//ff:what 핸들러 본문에서 form/json 바디 필드 추출을 검증한다
package flask

import (
	"strings"
	"testing"
)

func TestExtractBodyFields(t *testing.T) {
	src := []byte(`def handler():
    u = request.form['username']
    p = request.form.get('password')
    data = request.get_json()
    title = data['title']
    body = request.json['body']
    name = request.get_json()['name']
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDefs := findAllByType(root, "function_definition")
	if len(funcDefs) == 0 {
		t.Fatal("no function_definition")
	}

	bf := extractBodyFields(funcDefs[0], src)

	form := strings.Join(bf.formFields, ",")
	if !strings.Contains(form, "username") || !strings.Contains(form, "password") {
		t.Errorf("form fields = %v, want username,password", bf.formFields)
	}
	json := strings.Join(bf.jsonFields, ",")
	if !strings.Contains(json, "title") || !strings.Contains(json, "body") || !strings.Contains(json, "name") {
		t.Errorf("json fields = %v, want title,body,name", bf.jsonFields)
	}
	if !bf.hasJSON {
		t.Error("hasJSON = false, want true")
	}
}
