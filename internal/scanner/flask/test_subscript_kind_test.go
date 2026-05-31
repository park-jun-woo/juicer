//ff:func feature=scan type=test topic=flask control=sequence
//ff:what subscriptKind subscript 베이스를 form/json/jsonVar/미상으로 분류 테스트
package flask

import "testing"

func TestSubscriptKind(t *testing.T) {
	if n, s := firstSubscript(t, "x = request.form['a']\n"); subscriptKind(n, s, nil) != "form" {
		t.Error("form")
	}
	if n, s := firstSubscript(t, "x = request.get_json()['a']\n"); subscriptKind(n, s, nil) != "json" {
		t.Error("call json")
	}
	if n, s := firstSubscript(t, "x = data['a']\n"); subscriptKind(n, s, map[string]bool{"data": true}) != "json" {
		t.Error("identifier jsonVar")
	}
	if n, s := firstSubscript(t, "x = other['a']\n"); subscriptKind(n, s, nil) != "" {
		t.Error("unknown")
	}
}
