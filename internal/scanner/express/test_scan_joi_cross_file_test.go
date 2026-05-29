//ff:func feature=scan type=test control=sequence topic=express
//ff:what E2E 스캔 테스트: validate(authValidation.register) 크로스파일 Joi 검증 스키마
package express

import "testing"

func TestScan_JoiCrossFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "validations/auth.validation.ts", joiValidationFixture)
	writeFile(t, dir, "routes/auth.ts", joiRouteFixture)
	writeFile(t, dir, "app.ts", joiAppFixture)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	reg := findEndpoint(result.Endpoints, "POST", "/auth/register")
	if reg == nil || reg.Request == nil || reg.Request.Body == nil {
		t.Fatal("expected register endpoint with body")
	}
	if len(reg.Request.Body.Fields) != 2 {
		t.Fatalf("expected 2 body fields, got %d", len(reg.Request.Body.Fields))
	}
	if !fieldRequired(reg.Request.Body.Fields, "email") {
		t.Error("expected email field marked required via .required()")
	}

	ve := findEndpoint(result.Endpoints, "POST", "/auth/verify-email")
	if ve == nil || ve.Request == nil {
		t.Fatal("expected verify-email endpoint with request")
	}
	if len(ve.Request.Query) != 1 || ve.Request.Query[0].Name != "token" {
		t.Fatalf("expected token query param, got %+v", ve.Request.Query)
	}
}
