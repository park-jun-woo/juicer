//ff:func feature=scan type=test control=sequence topic=django
//ff:what 순수 Django 뷰 폴백이 GET 엔드포인트 1건과 path param을 생성하는지 검증한다
package django

import "testing"

func TestBuildPlainViewEndpoints(t *testing.T) {
	eps := buildPlainViewEndpoints(urlEntry{pattern: "items/<int:pk>/"}, "ItemView")
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	ep := eps[0]
	if ep.Method != "GET" || ep.Path != "/items/{pk}/" || ep.Handler != "ItemView" {
		t.Errorf("unexpected endpoint: %+v", ep)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "pk" {
		t.Errorf("expected pk path param, got %+v", ep.Request)
	}
}
