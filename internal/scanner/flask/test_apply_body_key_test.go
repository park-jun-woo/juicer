//ff:func feature=scan type=test topic=flask control=sequence
//ff:what applyBodyKey form/json 키 분류 기록 및 hasJSON 표시 테스트
package flask

import "testing"

func TestApplyBodyKey(t *testing.T) {
	bf := applyBodyKey(bodyFields{}, "form", "username")
	if len(bf.formFields) != 1 || bf.formFields[0] != "username" {
		t.Errorf("form: %+v", bf)
	}
	bf = applyBodyKey(bf, "json", "email")
	if !bf.hasJSON || len(bf.jsonFields) != 1 {
		t.Errorf("json: %+v", bf)
	}
	// json with empty key still marks hasJSON
	bf2 := applyBodyKey(bodyFields{}, "json", "")
	if !bf2.hasJSON || len(bf2.jsonFields) != 0 {
		t.Errorf("empty json key: %+v", bf2)
	}
	// unknown kind -> unchanged
	bf3 := applyBodyKey(bodyFields{}, "other", "x")
	if bf3.hasJSON || len(bf3.formFields) != 0 {
		t.Errorf("unknown: %+v", bf3)
	}
}
