//ff:func feature=scan type=test control=sequence topic=flask
//ff:what urlParamsToScannerParams 테스트
package flask

import "testing"

func TestURLParamsToScannerParams(t *testing.T) {
	params := []urlParam{
		{name: "user_id", converter: "int"},
		{name: "slug", converter: ""},
	}
	got := urlParamsToScannerParams(params)
	if len(got) != 2 {
		t.Fatalf("expected 2 params, got %d", len(got))
	}
	if got[0].Name != "user_id" || got[0].Type != "integer" {
		t.Errorf("param 0 = %+v", got[0])
	}
	if got[1].Name != "slug" || got[1].Type != "string" {
		t.Errorf("param 1 = %+v", got[1])
	}

	if len(urlParamsToScannerParams(nil)) != 0 {
		t.Error("expected empty for nil input")
	}
}
