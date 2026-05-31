//ff:func feature=scan type=test topic=flask control=sequence
//ff:what applyFormFields 폼 필드를 엔드포인트 요청에 반영(빈 키 무시) 테스트
package flask

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyFormFields(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyFormFields(ep, nil)
	if ep.Request != nil {
		t.Error("no keys -> no request")
	}
	applyFormFields(ep, []string{"username", "password"})
	if ep.Request == nil || len(ep.Request.FormFields) != 2 {
		t.Errorf("form fields: %+v", ep.Request)
	}
	if ep.Request.FormFields[0].Name != "username" {
		t.Errorf("name: %+v", ep.Request.FormFields)
	}
}
